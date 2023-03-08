package caches

import (
	"sync/atomic"
)

var (
	mid int32 = 0
)

func GetMid() int {
	addInt32 := atomic.AddInt32(&mid, 1)
	return int(addInt32 & 0xffff)
}
