package httpclient

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"github.com/nnkolesnikov/transport-learning/pkg/models"
	"github.com/nnkolesnikov/transport-learning/pkg/service"
	"github.com/nnkolesnikov/transport-learning/pkg/service/httpserver"
)

const (
	port                     = "8080"
	serverURL                = "http://localhost" + ":" + port
	serverHost               = "localhost" + ":" + port
	maxConns                 = 512
	maxRequestBodySize       = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second

	validId    = 10
	nonValidId = -1

	methodGetUser                = "GetUser"
	methodGetOrders              = "GetOrders"
	methodGetUserCount           = "GetUserCount"
	methodGetOrdersWithoutParams = "GetOrdersWithoutParams"

	nameTestClientGetUserSuccess                = "TestClientGetUserValidId"
	nameTestClientGetOrdersSuccess              = "TestClientGetOrdersValidId"
	nameTestClientGetUserCountSuccess           = "TestClientGetUserCountValidId"
	nameTestClientGetOrdersWithoutParamsSuccess = "TestClientGetOrdersWithoutParamsSuccess"

	nameTestClientGetUserFail      = "TestClientGetUserNonValidId"
	nameTestClientGetOrdersFail    = "TestClientGetOrdersNonValidId"
	nameTestClientGetUserCountFail = "TestClientGetUserCountNonValidId"
)

var (
	errNil      error
	errBadId    = errors.New("id <= 0")
	errInternal = service.NewError(http.StatusInternalServerError, "internal error")
)

func TestClient_GetUser_Success(t *testing.T) {
	t.Run(nameTestClientGetUserSuccess, func(t *testing.T) {
		request, response := makeGetUserRequest(validId), makeGoodDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetUser, context.Background(), request).Return(response, errNil)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		res, err := client.GetUser(context.Background(), request)
		assert.Equal(t, res, response)
		assert.NoError(t, err, "unexpected error:", err)
	})
}

func TestClient_GetOrders_Success(t *testing.T) {
	t.Run(nameTestClientGetOrdersSuccess, func(t *testing.T) {
		request, response := makeGetOrdersRequest(validId), makeGoodDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetOrders, context.Background(), request).Return(response, errNil)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		res, err := client.GetOrders(context.Background(), request)
		assert.Equal(t, res, response)
		assert.NoError(t, err, "unexpected error:", err)
	})
}

func TestClient_GetUserCount_Success(t *testing.T) {
	t.Run(nameTestClientGetUserCountSuccess, func(t *testing.T) {
		request, response := makeGetUserCountRequest(validId), makeGoodDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetUserCount, context.Background(), request).Return(response, errNil)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		res, err := client.GetUserCount(context.Background(), request)
		assert.Equal(t, res, response)
		assert.NoError(t, err, "unexpected error:", err)
	})
}

func TestClient_GetOrdersWithoutParams_Success(t *testing.T) {
	t.Run(nameTestClientGetOrdersWithoutParamsSuccess, func(t *testing.T) {
		response := models.DefaultResponse{}
		svcMock := new(service.MockService)
		svcMock.On(methodGetOrdersWithoutParams, context.Background()).Return(response, errNil)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		res, err := client.GetOrdersWithoutParams(context.Background())
		assert.Equal(t, res, response)
		assert.NoError(t, err, "unexpected error:", err)
	})
}

func TestClient_GetUser_Fail(t *testing.T) {
	t.Run(nameTestClientGetUserFail, func(t *testing.T) {
		request, response := makeGetUserRequest(nonValidId), makeBadDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetUser, context.Background(), request).Return(response, errBadId)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		_, err := client.GetUser(context.Background(), request)
		assert.Equal(t, err, errInternal)
	})
}

func TestClient_GetOrders_Fail(t *testing.T) {
	t.Run(nameTestClientGetOrdersFail, func(t *testing.T) {
		request, response := makeGetOrdersRequest(nonValidId), makeBadDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetOrders, context.Background(), request).Return(response, errBadId)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		_, err := client.GetOrders(context.Background(), request)
		assert.Equal(t, err, errInternal)
	})
}

func TestClient_GetUserCount_Fail(t *testing.T) {
	t.Run(nameTestClientGetUserCountFail, func(t *testing.T) {
		request, response := makeGetUserCountRequest(nonValidId), makeBadDefaultResponse()
		svcMock := new(service.MockService)
		svcMock.On(methodGetUserCount, context.Background(), request).Return(response, errBadId)
		server, client := makeServerClient(serverURL, svcMock)
		defer func() {
			err := server.Shutdown()
			if err != nil {
				log.Printf("server shut down err: %v", err)
			}
		}()
		time.Sleep(serverLaunchingWaitSleep)

		_, err := client.GetUserCount(context.Background(), request)
		assert.Equal(t, err, errInternal)
	})
}

func makeGetUserRequest(id int) *models.GetUserRequest {
	return &models.GetUserRequest{
		Id: id,
	}
}

func makeGetOrdersRequest(id int) *models.GetOrdersRequest {
	return &models.GetOrdersRequest{
		Id: id,
	}
}

func makeGetUserCountRequest(id int) *models.GetUserCountRequest {
	return &models.GetUserCountRequest{
		Id: id,
	}
}

func makeGoodDefaultResponse() models.DefaultResponse {
	return models.DefaultResponse{
		Data: &models.Data{Res: true},
	}
}

func makeBadDefaultResponse() models.DefaultResponse {
	return models.DefaultResponse{}
}

func makeServerClient(serverURL string, svc service.Service) (server *fasthttp.Server, client service.Service) {
	errorProcessor := service.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	client = NewPreparedClient(
		serverURL,
		serverHost,
		maxConns,
		nil,
		errorProcessor,
		service.NewError,
		URIPathClientGetUser,
		URIPathClientGetOrders,
		URIPathClientGetUserCount,
		URIPathClientGetOrdersWithoutParams,
		HTTPMethodGetUser,
		HTTPMethodGetOrders,
		HTTPMethodGetUserCount,
		HTTPMethodGetOrdersWithoutParams,
	)
	router := httpserver.NewPreparedServer(
		svc,
		errorProcessor,
		service.NewError,
	)
	server = &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: maxRequestBodySize,
		ReadTimeout:        serverTimeout,
	}
	go func() {
		err := server.ListenAndServe(serverHost)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}
