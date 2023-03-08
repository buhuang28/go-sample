package main

import "fmt"

type Stu struct {
	Name string
}

func main() {

	c := make(chan string, 10)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
