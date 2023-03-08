package main

import (
	"fmt"
	"time"
)

type Stu struct {
	Name string
}

func main() {
	AStu := Stu{Name: "A同学"}
	BStu := Stu{Name: "B同学"}

	ch := make(chan Stu, 1)

	go func() {
		for {
			select {
			case r := <-ch:
				fmt.Printf("%s,我要进来了\n", r.Name)
			case <-time.After(time.Second):
				fmt.Println("怎么还没人进来")
			}
		}
	}()
	fmt.Println("wait 2 second")
	time.Sleep(time.Second * 2)
	ch <- AStu
	ch <- BStu
	time.Sleep(time.Second)
}
