package main

import (
	"context"
	"fmt"
	"go-sample/go-grpc/pb/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Ask(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	fmt.Println("收到请求数据:", in.Name, in.Question)
	return &pb.HelloResp{
		Anwser: "What's your problem",
		To:     in.Name,
	}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":10088")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
