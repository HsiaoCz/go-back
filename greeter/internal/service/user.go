package service

import (
	"context"
	v1 "go-back/greeter/api/v1"
	"go-back/greeter/internal/biz"
)

type UserService struct {
	v1.UnimplementedGreeterServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Login(ctx context.Context, in *v1.UserLoginRequest) (*v1.UserLoginResponse, error) {
	return nil, nil
}

func (s *UserService) Singup(ctx context.Context, in *v1.UserSingupRequest) (*v1.UserSingupResponse, error) {
	return nil, nil
}
