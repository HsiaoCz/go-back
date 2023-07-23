package main

import (
	"context"
	"fmt"
	"go-back/grpc/interc/client/pb"
	"net"

	"google.golang.org/grpc"
)

// server
// 基于拦截器和metadata 实现token
type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func main() {
	lisn, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	// 启动服务
	err = s.Serve(lisn)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
