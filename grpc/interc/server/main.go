package main

import (
	"context"
	"fmt"
	"go-back/grpc/interc/client/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// server
// 基于拦截器和metadata 实现token
type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func InterToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return resp, status.Error(codes.Unauthenticated, "无token认真")
	}
	var (
		appid  string
		appkey string
	)

	if va1, ok := md["appid"]; ok {
		appid = va1[0]
	}
	if va2, ok := md["appkey"]; ok {
		appkey = va2[0]
	}
	if appid != "101010" || appkey != "this is key" {
		return resp, status.Error(codes.Unauthenticated, "认证错误")
	}
	res, err := handler(ctx, req)
	return res, err
}

func main() {
	lisn, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
		return
	}
	// 注册一个拦截器
	opt := grpc.UnaryInterceptor(InterToken)
	// 创建grpc服务
	s := grpc.NewServer(opt)
	// 在grpc服务端注册服务
	pb.RegisterHelloServer(s, &server{})
	// 启动服务
	err = s.Serve(lisn)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
