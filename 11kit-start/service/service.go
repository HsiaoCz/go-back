package service

import (
	"context"
	"errors"
)

// AddService 列出当前服务所有RPC方法的接口
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// addService 实现AddService接口
type addService struct{}

var (
	ErrEmptyString = errors.New("两个参数都是空字符串")
)

func (as *addService) Sum(_ context.Context, a, b int) (int, error) {
	// 业务逻辑
	return a + b, nil
}

func (as *addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}

	return a + b, nil
}

func NewService() AddService {
	return &addService{}
}
