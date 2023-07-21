package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service) //注册服务
	rpc.HandleHTTP()      // 基于HTTP协议
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}

