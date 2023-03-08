package pool

import (
	"go-sample/tcp_sample/tlv"
)

var (
	TLVMap = map[int]func() tlv.TLVer{
		tlv.LOGIN_TAG:   GetLoginPool,
		tlv.MESSAGE_TAG: GetMessagePool,
	}
)

func GetLoginPool() tlv.TLVer {
	return LoginPool.Get().(tlv.TLVer)
}

func GetMessagePool() tlv.TLVer {
	return MessagePool.Get().(tlv.TLVer)
}
