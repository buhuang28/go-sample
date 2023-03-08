package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	id int32 = 1
)

func main() {
	for i := 0; i < 75535; i++ {
		go fmt.Println(GetId())
	}
	time.Sleep(time.Second * 15)
}

func GetId() int32 {
	addInt32 := atomic.AddInt32(&id, 1)
	return addInt32 & 0xffff
}
