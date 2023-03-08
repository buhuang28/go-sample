package chs

import (
	"go-sample/tcp_sample/tlv"
	"sync"
)

var (
	tlvChs = make(map[int]chan tlv.TLVer)
	tLock  sync.RWMutex
)

func AddTLVChs(id int, t chan tlv.TLVer) {
	tLock.Lock()
	defer tLock.Unlock()
	tlvChs[id] = t
}

func GetTlvChs(id int) chan tlv.TLVer {
	tLock.RLock()
	defer tLock.RUnlock()
	return tlvChs[id]
}

func DelTlvChs(id int) {
	tLock.Lock()
	defer tLock.Unlock()
	delete(tlvChs, id)
}
