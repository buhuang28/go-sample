package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	xCtx := context.WithValue(ctx, "x", "a")
	vCtx := context.WithValue(xCtx, "v", "ssa")
	fmt.Println(vCtx.Value("x"))
	fmt.Println(vCtx.Value("v"))

	cancelCtx, cancelFunc := context.WithCancel(vCtx)

	go func() {
		time.Sleep(time.Second * 3)
		cancelFunc()
		fmt.Println("execute cancel")
	}()

	select {
	case <-time.After(time.Second):
		fmt.Println("run")
	case <-cancelCtx.Done():
		fmt.Println("done")
	}
	time.Sleep(time.Second * 5)
}
