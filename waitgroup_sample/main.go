package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	m := make(map[int]struct{})
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("print ", i)
			m[i] = struct{}{}
		}(i)
	}
	wg.Wait()
	fmt.Println("map len:", len(m))
}
