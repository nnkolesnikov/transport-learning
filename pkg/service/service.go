package service

import (
	"context"
	"errors"

	"github.com/nnkolesnikov/transport-learning/pkg/models"
)

// Service ...
type Service interface {
	GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error)
	GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error)
	GetOrdersWithoutParams(ctx context.Context) (response models.DefaultResponse, err error)
}

type service struct{}

func (s *service) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error) {
	if request.Id > 0 {
		response.Data.Res = true
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = errors.New("id <= 0")
	return
}

func (s *service) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error) {
	if request.Id > 0 {
		response.Data.Res = true
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = errors.New("id <= 0")
	return
}

func (s *service) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error) {
	if request.Id > 0 {
		response.Data.Res = true
		return
	}
	response.Error = true
	response.ErrorText = "bad id"
	err = errors.New("id <= 0")
	return
}

func (s *service) GetOrdersWithoutParams(ctx context.Context) (response models.DefaultResponse, err error) {
	return
}

// NewService ...
func NewService() (Service, error) {
	return &service{}, nil
}
