package main

import (
	"context"
	"errors"
	"fmt"
	"go-back/grpc/grpc/server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// hello server

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal(errors.New("failed to get metadata"))
	}
	// 将metadata打印出来
	// 这里有一点，md是一个map[string]string结构
	fmt.Println(md["hello"][0])
	return &pb.HelloResponse{Reply: "hello" + in.Name}, nil
}

func main() {
	// 监听本地的9091端口
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
		return
	}

	s := grpc.NewServer()                  // 创建grpc服务
	pb.RegisterGreeterServer(s, &server{}) // 在grpc服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
