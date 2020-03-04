package httpclient

import (
	"github.com/nnkolesnikov/transport-learning/pkg/service"
	"github.com/nnkolesnikov/transport-learning/pkg/service/httpserver"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	Port                     = "8080"
	serverURL                = "http://localhost" + ":" + Port
	serverHost               = "localhost" + ":" + Port
	maxConns                 = 512
	maxRequestBodySize       = 15 * 1024 * 1024
	serverTimeout            = 1 * time.Millisecond
	serverLaunchingWaitSleep = 1 * time.Second
)

func TestClient_GetUser_Success(t *testing.T) {

}

func TestClient_GetOrders_Success(t *testing.T) {

}

func TestClient_GetUserCount_Success(t *testing.T) {

}

func TestClient_GetOrdersWithoutParams_Success(t *testing.T) {

}

func TestClient_GetUser_Fail(t *testing.T) {

}

func TestClient_GetOrders_Fail(t *testing.T) {

}

func TestClient_GetUserCount_Fail(t *testing.T) {

}

func TestClient_GetOrdersWithoutParams_Fail(t *testing.T) {

}

func makeServerClient(serverURL string, svc Service) (server *fasthttp.Server, client Service) {
	errorProcessor := service.NewErrorProcessor(http.StatusInternalServerError, "iternal error")
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
		err := server.ListenAndServe(serverURL)
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()

	return
}
