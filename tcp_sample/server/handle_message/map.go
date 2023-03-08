package handle_message

import (
	log "github.com/sirupsen/logrus"
	"go-sample/tcp_sample/tlv"
	"net"
)

var (
	handleMap = map[int]func(tlv.TLVer) ([]byte, error){
		tlv.LOGIN_TAG:   HandleLogin,
		tlv.MESSAGE_TAG: HandleMessage,
	}
)

func HandleConn(conn net.Conn, tag int, t tlv.TLVer) {
	handleFun, ok := handleMap[tag]
	if !ok {
		log.Error("not found function for", tag)
		_ = conn.Close()
		return
	}
	result, err := handleFun(t)
	if err != nil {
		log.Error(err)
		_ = conn.Close()
		return
	}
	if len(result) == 0 {
		log.Error("result is null")
		_ = conn.Close()
		return
	}
	_, err = conn.Write(result)
	if err != nil {
		log.Error(err)
		_ = conn.Close()
		return
	}
}
