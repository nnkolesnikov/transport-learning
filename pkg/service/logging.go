//Package service logging wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (s *loggingMiddleware) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetUser",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetUser(ctx, request)
}

func (s *loggingMiddleware) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetOrders",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetOrders(ctx, request)
}

func (s *loggingMiddleware) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetUserCount",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetUserCount(ctx, request)
}

func (s *loggingMiddleware) GetOrdersWithoutParams(ctx context.Context,  ) (response models.DefaultResponse, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetOrdersWithoutParams",
			
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetOrdersWithoutParams(ctx, )
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
