package main

import (
	"fmt"
	"sync"
)

type People interface {
	Add(s string)
	Run()
}

type Stu struct {
	Name string
}

func (s *Stu) Add(a string) {
	s.Name = a
}

func (s *Stu) Run() {
	fmt.Println(s.Name, "run")
}

var StuPool = sync.Pool{
	New: func() any {
		return new(Stu)
	},
}

func main() {
	peo := StuPool.Get().(People)
	peo.Add("lalla")
	PeopleRun(peo)
	peo2 := StuPool.Get().(People)
	PeopleRun(peo2)
}

func PeopleRun(p People) {
	p.Run()
	StuPool.Put(p)
}
