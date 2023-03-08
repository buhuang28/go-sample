package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	Name   string
	Age    int
	Status bool
}

func init() {
	err := initClient(RedisConfig{
		Ip:       "127.0.0.1",
		Port:     "6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSetAndGet(t *testing.T) {
	var err error
	a := A{
		Name:   "asdsad",
		Age:    23,
		Status: true,
	}
	marshal, _ := json.Marshal(a)
	err = SetVal("afsdf", marshal, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = SetVal("123213x", 12.34, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	val := GetVal("123213")
	fmt.Println(val)
	err = PutList_R("oks", 55466)
	if err != nil {
		fmt.Println(err)
		return
	}
	val = GetVal("oks")
	fmt.Println(val)
}

func TestGetAndGetList(t *testing.T) {
	err := PutList_R("myksey", "asdasdasdasd")
	if err != nil {
		fmt.Println(err)
		return
	}

	err, strings := RangeList_R("myksey")
	fmt.Println(strings)
	fmt.Println(err)
}

func TestOa(t *testing.T) {
	a := 0b00000010
	fmt.Println(a & 0b10 >> 1)
}
