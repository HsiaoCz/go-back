package main

import (
	"context"
	"fmt"
	"go-back/grpc/interc/client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
// 基于拦截器和metadata 实现token
func main() {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
