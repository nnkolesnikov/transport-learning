//Package service http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"net/http"

	"github.com/valyala/fasthttp"
)

//type errorProcessor interface {
//	Encode(ctx context.Context, r *fasthttp.Response, err error)
//}

type service interface {
	GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error)
	GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error)
	GetOrdersWithoutParams(ctx context.Context,  ) (response models.DefaultResponse, err error)
}


type getUserServer struct {
	transport      GetUserTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getUserServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetUser(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetUserServer the server creator
func NewGetUserServer(transport GetUserTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getUserServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getOrdersServer struct {
	transport      GetOrdersTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getOrdersServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetOrders(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetOrdersServer the server creator
func NewGetOrdersServer(transport GetOrdersTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getOrdersServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getUserCountServer struct {
	transport      GetUserCountTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getUserCountServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetUserCount(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetUserCountServer the server creator
func NewGetUserCountServer(transport GetUserCountTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getUserCountServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getOrdersWithoutParamsServer struct {
	transport      GetOrdersWithoutParamsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getOrdersWithoutParamsServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetOrdersWithoutParams(ctx, &)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetOrdersWithoutParamsServer the server creator
func NewGetOrdersWithoutParamsServer(transport GetOrdersWithoutParamsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getOrdersWithoutParamsServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

func NewPreparedServer(svc service) {
	errorProcessor := NewErrorProcessor(http.StatusInternalServerError, "internal error")
	getOrdersTransport := NewGetOrdersTransport()
	getUserCountTransport := NewGetUserCountTransport()	
	getUserTransport := NewGetUserTransport()
	getOrdersWithoutParamsTransport := NewGetOrdersWithoutParamsTransport() 
	return MakeFastHTTPRouter(
		[]*barcodeService.HandlerSettings{
			{
				Path:    URIPathClientGetUser,
				Method:  http.MethodGet,
				Handler: NewGetUserServer(getUserTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathGetOrders,
				Method:  http.MethodPost,
				Handler: NewGetOrdersServer(getOrdersTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathClientGetUserCount,
				Method:  http.MethodGet,
				Handler: NewGetUserCountServer(getUserCountTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathGetOrdersWithoutParams,
				Method:  http.MethodGet,
				Handler: NewGetOrdersWithoutParamsServer(getOrdersWithoutParamsTransport, svc, errorProcessor),
			},
		},
	)
}
