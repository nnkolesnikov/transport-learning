package httpclient

import (
	"context"
	"github.com/nnkolesnikov/transport-learning/pkg/models"

	"github.com/valyala/fasthttp"
)

var (
	GetUser                = option{}
	GetOrders              = option{}
	GetUserCount           = option{}
	GetOrdersWithoutParams = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {
	GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error)
	GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error)
	GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error)
	GetOrdersWithoutParams(ctx context.Context, ) (response models.DefaultResponse, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetUser                GetUserClientTransport
	transportGetOrders              GetOrdersClientTransport
	transportGetUserCount           GetUserCountClientTransport
	transportGetOrdersWithoutParams GetOrdersWithoutParamsClientTransport
	options                         map[interface{}]Option
}

// GetUser ...
func (s *client) GetUser(ctx context.Context, request *models.GetUserRequest) (response models.DefaultResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetUser]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUser.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetUser.DecodeResponse(ctx, res)
}

// GetOrders ...
func (s *client) GetOrders(ctx context.Context, request *models.GetOrdersRequest) (response models.DefaultResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetOrders]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrders.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetOrders.DecodeResponse(ctx, res)
}

// GetUserCount ...
func (s *client) GetUserCount(ctx context.Context, request *models.GetUserCountRequest) (response models.DefaultResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetUserCount]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetUserCount.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetUserCount.DecodeResponse(ctx, res)
}

// GetOrdersWithoutParams ...
func (s *client) GetOrdersWithoutParams(ctx context.Context) (response models.DefaultResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetOrdersWithoutParams]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrdersWithoutParams.EncodeRequest(ctx, req, ); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetOrdersWithoutParams.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportGetUser GetUserClientTransport,
	transportGetOrders GetOrdersClientTransport,
	transportGetUserCount GetUserCountClientTransport,
	transportGetOrdersWithoutParams GetOrdersWithoutParamsClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli:                             cli,
		transportGetUser:                transportGetUser,
		transportGetOrders:              transportGetOrders,
		transportGetUserCount:           transportGetUserCount,
		transportGetOrdersWithoutParams: transportGetOrdersWithoutParams,
		options:                         options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	options map[interface{}]Option,
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathGetUser string,
	uriPathGetOrders string,
	uriPathGetUserCount string,
	uriPathGetOrdersWithoutParams string,

	httpMethodGetUser string,
	httpMethodGetOrders string,
	httpMethodGetUserCount string,
	httpMethodGetOrdersWithoutParams string,
) Service {
	transportGetUser := NewGetUserClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUser,
		httpMethodGetUser,
	)

	transportGetOrders := NewGetOrdersClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrders,
		httpMethodGetOrders,
	)

	transportGetUserCount := NewGetUserCountClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetUserCount,
		httpMethodGetUserCount,
	)

	transportGetOrdersWithoutParams := NewGetOrdersWithoutParamsClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetOrdersWithoutParams,
		httpMethodGetOrdersWithoutParams,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},
		transportGetUser,
		transportGetOrders,
		transportGetUserCount,
		transportGetOrdersWithoutParams,
		options,
	)
}
