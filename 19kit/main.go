package main

import (
	"context"
	"go-back/19kit/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type HelloService interface {
	SayHello(context.Context, string) (string, error)
}

type grpcServer struct {
	pb.UnimplementedHelloServer
	sayHello grpctransport.Handler
}

type HelloResponse struct {
	Msg string `json:"msg"`
}

type HelloRequest struct {
	Name string `json:"name"`
}

func makeSayHelloEndpoint(srv HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HelloRequest)
		res, err := srv.SayHello(ctx, req.Name)
		if err != nil {
			return HelloResponse{Msg: err.Error()}, nil
		}
		return HelloResponse{Msg: res}, nil
	}
}

func decodeGRPCSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HelloRequest)
	return HelloRequest{Name: req.Name}, nil
}

func encodeGRPCSayHelloResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*pb.HelloResponse)
	return HelloResponse{Msg: res.Msg}, nil
}

func (s *grpcServer) SayHello(ctx context.Context, req string) (string, error) {
	return "hi" + req, nil
}
