package endpoint

import (
	"context"
	"go-back/11kit-start/service"

	"github.com/go-kit/kit/endpoint"
)

type SumRequset struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func MakeSumEndpoint(srv service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SumRequset)
		v, err := srv.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func MakeConcatEnpoint(srv service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ConcatRequest)
		v, err := srv.Concat(ctx, req.A, req.B)
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}
