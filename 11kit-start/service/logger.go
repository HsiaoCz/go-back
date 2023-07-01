package service

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type logMiddleware struct {
	logger log.Logger
	next   AddService
}

func (lm logMiddleware) Sum(ctx context.Context, a int, b int) (res int, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "sum",
			"a", a,
			"b", b,
			"output", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = lm.next.Sum(ctx, a, b)
	return
}

func (lm logMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"method", "sum",
			"a", a,
			"b", b,
			"output", res,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	res, err = lm.next.Concat(ctx, a, b)
	return
}

// 创建一个带日志的add service
func NewLogMiddleware(logger log.Logger, svc AddService) AddService {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}
