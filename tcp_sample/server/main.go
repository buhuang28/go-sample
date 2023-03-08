package main

import "go-sample/tcp_sample/server/log_init"

func init() {
	log_init.InitLog("./", 7)
}

func main() {
	//多网卡的话需要指定IP
	_ = NewServer("", "7535")
	select {}
}
