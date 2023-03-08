package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-sample/tcp_sample/server/handle_message"
	"go-sample/tcp_sample/server/pool"
	"go-sample/utils"
	"io"
	"net"
)

func NewServer(Ip, Port string) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", Ip, Port))
	if err != nil {
		panic(err)
	}
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				log.Error(e)
			}
		}()
		for {
			conn, acErr := listen.Accept()
			if acErr != nil {
				log.Error(acErr)
				continue
			}
			go HandleConn(conn)
		}
	}()
	return err
}

const (
	TAG_LEN = 2
)

func HandleConn(conn net.Conn) {
	defer func() {
		e := recover()
		if e != nil {
			log.Error(e)
		}
		conn.Close()
	}()

	bufRead := bufio.NewReader(conn)
	for {
		tag := pool.TagPool.Get().([]byte)
		_, err := io.ReadAtLeast(bufRead, tag, TAG_LEN)
		if err != nil {
			log.Error(err)
			return
		}
		getTLV, ok := pool.TLVMap[utils.LitBytes2Int(tag)]
		if !ok {
			log.Error("unknow tag:", tag)
			return
		}
		t := getTLV()
		h := make([]byte, t.GetHeaderLen())
		_, err = io.ReadAtLeast(bufRead, h, t.GetHeaderLen())
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
		_, err = io.ReadAtLeast(bufRead, v, t.GetBodyLen())
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
			handle_message.HandleConn(conn, utils.LitBytes2Int(tag), t)
		}()
	}

}
