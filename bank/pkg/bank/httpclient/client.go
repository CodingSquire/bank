package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"github.com/CodingSquire/bank/pkg/api"
)

// Service implements Service interface
type Service interface {
	GetBalance(request *api.GetBalanceRequest) (response api.GetBalanceResponse, err error)
}

type fasthttpSS interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
}

type client struct {
	cli fasthttpSS

	transportGetBalance GetBalanceTransport
}

// GetBrandsByID ...
func (s *client) GetBalance(request *api.GetBalanceRequest) (response api.GetBalanceResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	ctx := context.Background()
	if err = s.transportGetBalance.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetBalance.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli fasthttpSS,
	transportGetBalance GetBalanceTransport,
) Service {
	return &client{
		cli: cli,

		transportGetBalance: transportGetBalance,
	}
}
