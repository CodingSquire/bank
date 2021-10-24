package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/CodingSquire/bank/pkg/api"
)

type errorCreator func(status int, format string, v ...interface{}) error

// GetBalanceTransport transport interface
type GetBalanceTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetBalanceRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetBalanceResponse) (err error)
}

type getBalanceTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getBalanceTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetBalanceRequest, err error) {
	err = json.Unmarshal(r.Body(), &request)
	if err != nil {
		return request, t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getBalanceTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetBalanceResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	if _, err = fasthttp.WriteBrotli(r.BodyWriter(), body); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return

}

// NewGetBalanceTransport the transport creator for http requests
func NewGetBalanceTransport(errorCreator errorCreator) GetBalanceTransport {
	return &getBalanceTransport{
		errorCreator: errorCreator,
	}
}

// CreateAccountTransport transport interface
type CreateAccountTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateAccRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateAccResponse) (err error)
}

type createAccountTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *createAccountTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateAccRequest, err error) {
	err = json.Unmarshal(r.Body(), &request)
	if err != nil {
		return request, t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createAccountTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateAccResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	if _, err = fasthttp.WriteBrotli(r.BodyWriter(), body); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewCreateAccountTransport the transport creator for http requests
func NewCreateAccountTransport(errorCreator errorCreator) CreateAccountTransport {
	return &createAccountTransport{
		errorCreator: errorCreator,
	}
}

// CreateAccountTransport transport interface
type AddToBellTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.AddToBellRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.AddToBellResponse) (err error)
}

type addToBellTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *addToBellTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.AddToBellRequest, err error) {
	err = json.Unmarshal(r.Body(), &request)
	if err != nil {
		return request, t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *addToBellTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.AddToBellResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	if _, err = fasthttp.WriteBrotli(r.BodyWriter(), body); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewCreateAccountTransport the transport creator for http requests
func NewAddToBellTransport(errorCreator errorCreator) AddToBellTransport {
	return &addToBellTransport{
		errorCreator: errorCreator,
	}
}
