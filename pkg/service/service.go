package service

import (
	"context"

	"github.com/nnkolesnikov/transport-learning/pkg/models"
)

// Service ...
type Service interface {
	GetUser(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetOrders(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetUserCount(ctx context.Context, request *models.Request) (response models.Response, err error)
	GetOrdersWithoutParams(ctx context.Context) (response models.Response, err error)
}

type service struct{}

func (s *service) GetUser(ctx context.Context, request *models.Request) (response models.Response, err error) {

	return
}

func (s *service) GetOrders(ctx context.Context, request *models.Request) (response models.Response, err error) {

	return
}

func (s *service) GetUserCount(ctx context.Context, request *models.Request) (response models.Response, err error) {

	return
}

func (s *service) GetOrdersWithoutParams(ctx context.Context) (response models.Response, err error) {

	return
}

// NewService ...
func NewService() (Service, error) {
	return &service{}, nil
}
