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

loop:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("run")
			goto loop
		case <-cancelCtx.Done():
			fmt.Println("done")
			break loop
		}
	}

	time.Sleep(time.Second * 5)
	fmt.Println("success")
}
