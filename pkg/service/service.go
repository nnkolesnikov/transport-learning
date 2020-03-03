package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/nnkolesnikov/transport-learning/pkg/models"
)

// Service ...
type Service interface {
	GetUser(ctx context.Context, request *models.GetUserRequest) (response models.Response, err error)
	GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.Response, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.Response, err error)
	GetOrdersWithoutParams(ctx context.Context) (response models.Response, err error)
}

type service struct{}

func (s *service) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.Response, err error) {
	if request.Id > 0 {
		response.Data.Res = 1
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = fmt.Errorf("bad id")
	return
}

func (s *service) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.Response, err error) {
	if request.Id > 0 {
		response.Data.Res = 1
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = errors.New("bad id")
	return
}

func (s *service) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.Response, err error) {
	if request.Id > 0 {
		response.Data.Res = 1
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = fmt.Errorf("bad id")
	return
}

func (s *service) GetOrdersWithoutParams(ctx context.Context) (response models.Response, err error) {
	return
}

// NewService ...
func NewService() (Service, error) {
	return &service{}, nil
}
