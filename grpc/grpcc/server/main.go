package main

import (
	"context"
	"fmt"
	"go-back/grpc/grpcc/server/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(crx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "hi" + in.Name}, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Msg: "login successed"}, nil
}

func (s *server) Singup(ctx context.Context, in *pb.SingupRequest) (*pb.SingupResponse, error) {
	return &pb.SingupResponse{Msg: "注册成功"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
