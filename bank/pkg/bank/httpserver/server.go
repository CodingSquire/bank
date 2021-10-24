package httpserver

import (
	"context"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/CodingSquire/bank/pkg/api"
	"github.com/CodingSquire/bank/pkg/httpserver"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {
	GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error)
	CreateAcc(request api.CreateAccRequest) (response api.CreateAccResponse, err error)
	AddToBell(request api.AddToBellRequest) (response api.AddToBellResponse, err error)
}

type getBalanceServer struct {
	transport      GetBalanceTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getBalanceServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetBalance(request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetBalanceServer the server creator
func NewGetBalanceServer(transport GetBalanceTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getBalanceServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type createAccountServer struct {
	transport      CreateAccountTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createAccountServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.CreateAcc(request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateAccountServer the server creator
func NewCreateAccountServer(transport CreateAccountTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createAccountServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type addToBellServer struct {
	transport      AddToBellTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *addToBellServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.AddToBell(request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewAddToBellServer the server creator
func NewAddToBellServer(transport AddToBellTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := addToBellServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	getBalanceTransport := NewGetBalanceTransport(httpserver.NewError)
	createAccountTransport := NewCreateAccountTransport(httpserver.NewError)
	addToBellTransport := NewAddToBellTransport(httpserver.NewError)

	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientGetBalance,
				Method: HTTPMethodGetBalance,
				Handler: NewGetBalanceServer(
					getBalanceTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientCreateAccount,
				Method: HTTPMethodCreateAccount,
				Handler: NewCreateAccountServer(
					createAccountTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientAddToBell,
				Method: HTTPMethodAddToBell,
				Handler: NewAddToBellServer(
					addToBellTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
