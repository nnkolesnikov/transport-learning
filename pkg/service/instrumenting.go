//Package service instrumenting wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

func (s *instrumentingMiddleware) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error) {
	defer s.recordMetrics("GetUser", time.Now(), err)
	return s.svc.GetUser(ctx, request)
}

func (s *instrumentingMiddleware) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error) {
	defer s.recordMetrics("GetOrders", time.Now(), err)
	return s.svc.GetOrders(ctx, request)
}

func (s *instrumentingMiddleware) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error) {
	defer s.recordMetrics("GetUserCount", time.Now(), err)
	return s.svc.GetUserCount(ctx, request)
}

func (s *instrumentingMiddleware) GetOrdersWithoutParams(ctx context.Context,  ) (response models.DefaultResponse, err error) {
	defer s.recordMetrics("GetOrdersWithoutParams", time.Now(), err)
	return s.svc.GetOrdersWithoutParams(ctx, )
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
