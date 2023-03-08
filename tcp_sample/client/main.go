package main

import (
	log "github.com/sirupsen/logrus"
	"go-sample/tcp_sample/client/caches"
	"go-sample/tcp_sample/client/log_init"
	"go-sample/tcp_sample/client/tcpclient"
	"go-sample/tcp_sample/cst"
	"go-sample/tcp_sample/pb/pb"
	"go-sample/tcp_sample/tlv"
	"google.golang.org/protobuf/proto"
	"time"
)

func init() {
	log_init.InitLog("./", 7)
}

func main() {
	err := tcpclient.NewClient("127.0.0.1", "7535", 30)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	log.Info("链接时间:", time.Unix(tcpclient.Cli.ConnectTime, 0).Format(cst.TIME_FORMAT))
	//第一次请求
	loginTLV := tlv.NewLoginTLV(caches.GetMid(), time.Now().Unix())
	loginReq := &pb.LoginReq{
		Account:  "asdf",
		Password: "asdasd",
	}
	marshal, _ := proto.Marshal(loginReq)
	loginTLV.AddVal(marshal)
	wait := tcpclient.SendAndWait(loginTLV)
	if wait == nil {
		log.Info("wait null data")
		return
	}
	loginTLVResp := wait.(*tlv.LoginTLV)
	resp := new(pb.LoginResp)
	_ = proto.Unmarshal(loginTLVResp.Val, resp)
	log.Info(resp.Result)
	log.Info(resp.Status)

	//第二次请求
	loginTLV = tlv.NewLoginTLV(caches.GetMid(), time.Now().Unix())
	loginReq.Account = "root"
	loginReq.Password = "123456"
	marshal, _ = proto.Marshal(loginReq)
	loginTLV.AddVal(marshal)
	wait = tcpclient.SendAndWait(loginTLV)
	if wait == nil {
		log.Info("wait null data")
		return
	}
	loginTLVResp = wait.(*tlv.LoginTLV)
	_ = proto.Unmarshal(loginTLVResp.Val, resp)
	log.Info(resp.Result)
	log.Info(resp.Status)

	//第三次请求
	messageTLV := tlv.NewMessageTLV(caches.GetMid(), time.Now().Unix(), 0, 0)
	messageReq := &pb.MessageReq{
		Data: "Hello,I am from asdasdsadasfdgdgffgh",
	}
	marshal, _ = proto.Marshal(messageReq)
	messageTLV.AddVal(marshal)
	wait = tcpclient.SendAndWait(messageTLV)
	if wait == nil {
		log.Info("wait null data")
		return
	}
	msgTLVResp := wait.(*tlv.MessageTLV)
	msgResp := new(pb.MessageResp)
	_ = proto.Unmarshal(msgTLVResp.Val, msgResp)
	log.Info(msgResp.Message)
	log.Info(msgResp.Status)
	log.Info(msgResp.Result)
}
