package main

import (
	"flag"
	"fmt"
	"go-back/11kit-start/pb"
	"go-back/11kit-start/service"
	"go-back/11kit-start/transport"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "HTTP listen address")
	grpcAddr = flag.String("grpc-addr", ":8972", "gRPC listen address")
)

func main() {
	bs := service.NewService()

	var g errgroup.Group

	// HTTP服务
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := transport.NewHTTPServer(bs)
		return http.Serve(httpListener, httpHandler)
	})

	g.Go(func() error {
		// gRPC服务
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			fmt.Printf("grpc: net.Listen(tcp, %s) faield, err:%v\n", *grpcAddr, err)
			return err
		}
		defer grpcListener.Close()
		s := grpc.NewServer()
		pb.RegisterAddServer(s, transport.NewGRPCServer(bs))
		return s.Serve(grpcListener)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with err:%v\n", err)
	}
}
