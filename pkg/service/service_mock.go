package service

import (
	"context"
	"github.com/nnkolesnikov/transport-learning/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error) {
	args := m.Called(context.Background(), request)

	if a, ok := args.Get(0).(models.DefaultResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(0)
}

func (m *MockService) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error) {
	args := m.Called(context.Background(), request)
	if a, ok := args.Get(0).(models.DefaultResponse); ok {
		return a, args.Error(1)
	}

	return response, args.Error(1)
}

func (m *MockService) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error) {
	args := m.Called(context.Background(), request)

	if a, ok := args.Get(0).(models.DefaultResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(1)
}

func (m *MockService) GetOrdersWithoutParams(ctx context.Context) (response models.DefaultResponse, err error) {
	args := m.Called(context.Background())

	if a, ok := args.Get(0).(models.DefaultResponse); ok {
		return a, args.Error(1)
	}
	return response, args.Error(1)
}
