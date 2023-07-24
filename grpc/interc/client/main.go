package main

import (
	"context"
	"fmt"
	"go-back/grpc/interc/client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// 基于拦截器和metadata 实现token

func InterToken(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{
		"addid":  "101010",
		"appkey": "this is key",
	})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func main() {
	// 注册一个拦截器
	opt := grpc.WithUnaryInterceptor(InterToken)
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{Username: "zhangsan", Password: "admin"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("reply:", r)
}
