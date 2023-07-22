package service

import (
	"context"
	"go-back/grpc/user/pb"
)

type UserService struct {
	pb.UnimplementedGreeterServer
}

func (u *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}
