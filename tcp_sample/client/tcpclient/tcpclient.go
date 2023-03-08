package tcpclient

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-sample/tcp_sample/client/chs"
	"go-sample/tcp_sample/client/pool"
	"go-sample/tcp_sample/tlv"
	"go-sample/utils"
	"io"
	"net"
	"time"
)

var (
	Cli *TCPClient
)

type TCPClient struct {
	Conn        net.Conn
	ConnectTime int64
	Ip          string
	Port        string
	Timeout     int64
}

func NewClient(ip, port string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, port), time.Second*timeout)
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	cli := &TCPClient{
		Conn:        conn,
		ConnectTime: now,
		Ip:          ip,
		Port:        port,
		Timeout:     now,
	}
	Cli = cli
	go func() {
		HandleConn(cli)
	}()
	return nil
}

func HandleConn(c *TCPClient) {
	reader := bufio.NewReader(c.Conn)
	defer func() {
		e := recover()
		if e != nil {
			log.Error(utils.PrintStackTrace(e))
		}
		_ = c.Conn.Close()
	}()
	for {
		tag := make([]byte, 2)
		_, err := io.ReadAtLeast(reader, tag, 2)
		if err != nil {
			log.Error()
			return
		}
		tlvMap, ok := pool.TLVMap[utils.LitBytes2Int(tag)]
		if !ok {
			log.Error("非法tag:", tag)
			return
		}
		t := tlvMap()
		h := make([]byte, t.GetHeaderLen())
		_, err = io.ReadAtLeast(reader, h, t.GetHeaderLen())
		if err != nil {
			log.Error(err)
			return
		}
		err = t.Bytes2Header(h)
		if err != nil {
			log.Error(err)
			return
		}
		v := make([]byte, t.GetBodyLen())
		_, err = io.ReadAtLeast(reader, v, t.GetBodyLen())
		if err != nil {
			log.Error(err)
			return
		}
		go func() {
			err = t.Bytes2Body(v)
			if err != nil {
				log.Error(err)
				return
			}
			ch := chs.GetTlvChs(utils.LitBytes2Int(h[:2]))
			if ch == nil {
				log.Error("未找到发送的消息id:", h[:2])
				return
			}
			ch <- t
			close(ch)
		}()
	}
}

func SendAndWait(t tlv.TLVer) tlv.TLVer {
	mid := t.GetMid()
	ch := make(chan tlv.TLVer, 1)
	chs.AddTLVChs(mid, ch)
	defer chs.DelTlvChs(mid)
	_, err := Cli.Conn.Write(t.ToBytes())
	if err != nil {
		log.Errorf("write error:%v", err)
		return nil
	}
	select {
	case <-time.After(time.Second * 1500):
		log.Info("wait timeout")
		return nil
	case r := <-ch:
		return r
	}
}
