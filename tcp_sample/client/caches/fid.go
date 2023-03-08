package caches

import "sync/atomic"

var (
	fid int32
)

func GetFid() int {
	addInt32 := atomic.AddInt32(&fid, 1)
	return int(addInt32 & 0xff)
}
