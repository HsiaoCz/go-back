package main

import (
	"context"
	"flag"
	"fmt"
	"go-back/grpc/grpc/client/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// hello client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:9091", "the address to connect to")
	name = flag.String("name", defaultName, "Name  to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// 执行rpc调用并打印收到的响应数据
	// 生成metadata
	md := metadata.New(map[string]string{"hello": "hi"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Greetin:%s", r.GetReply())
}
