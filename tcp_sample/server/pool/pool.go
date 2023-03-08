package pool

import (
	"go-sample/tcp_sample/tlv"
	"sync"
)

var LoginPool = sync.Pool{
	New: func() any {
		return new(tlv.LoginTLV)
	},
}

var MessagePool = sync.Pool{
	New: func() any {
		return new(tlv.MessageTLV)
	},
}

var TagPool = sync.Pool{
	New: func() any {
		return make([]byte, 2)
	},
}
