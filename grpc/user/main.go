package main

import (
	"context"
	"go-back/grpc/user/pb"
)

type UserCase struct {
	pb.UnimplementedGreeterServer
}

func (u *UserCase) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func (u *UserCase) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}
