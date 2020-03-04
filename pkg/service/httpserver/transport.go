package httpserver

import (
	"context"
	"github.com/nnkolesnikov/transport-learning/pkg/models"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
)

// GetUserTransport transport interface
type GetUserTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error)
}

type getUserTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetUserRequest, err error) {
	request.Id, err = strconv.Atoi(string(r.URI().QueryArgs().Peek("id")))
	if err != nil {
		return request, t.errorCreator(
			http.StatusBadRequest,
			"Bad request, check the fields.",
			"failed to get Id from query: %v",
			err,
		)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(
			http.StatusInternalServerError,
			"failed to encode JSON response: %s",
			err,
		)
	}
	return
}

// NewGetUserTransport the transport creator for http requests
func NewGetUserTransport(errorCreator errorCreator) GetUserTransport {
	return &getUserTransport{
		errorCreator: errorCreator,
	}
}

// GetOrdersTransport transport interface
type GetOrdersTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetOrdersRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error)
}

type getOrdersTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getOrdersTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request models.GetOrdersRequest, err error) {
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		return models.GetOrdersRequest{}, t.errorCreator(http.StatusBadRequest, "failed to decode JSON request: %v", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getOrdersTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(
			http.StatusInternalServerError,
			"failed to encode JSON response: %s",
			err,
		)
	}
	return
}

// NewGetOrdersTransport the transport creator for http requests
func NewGetOrdersTransport(errorCreator errorCreator) GetOrdersTransport {
	return &getOrdersTransport{
		errorCreator: errorCreator,
	}
}

// GetUserCountTransport transport interface
type GetUserCountTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx) (request models.GetUserCountRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error)
}

type getUserCountTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getUserCountTransport) DecodeRequest(ctx *fasthttp.RequestCtx) (request models.GetUserCountRequest, err error) {
	request.Id, err = strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		return request, t.errorCreator(
			http.StatusBadRequest,
			"Bad request, check the fields.",
			"failed to get Id from URI: %v",
			err,
		)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserCountTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetUserCountTransport the transport creator for http requests
func NewGetUserCountTransport(errorCreator errorCreator) GetUserCountTransport {
	return &getUserCountTransport{
		errorCreator: errorCreator,
	}
}

// GetOrdersWithoutParamsTransport transport interface
type GetOrdersWithoutParamsTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error)
}

type getOrdersWithoutParamsTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getOrdersWithoutParamsTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getOrdersWithoutParamsTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *models.DefaultResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetOrdersWithoutParamsTransport the transport creator for http requests
func NewGetOrdersWithoutParamsTransport(errorCreator errorCreator) GetOrdersWithoutParamsTransport {
	return &getOrdersWithoutParamsTransport{
		errorCreator: errorCreator,
	}
}
