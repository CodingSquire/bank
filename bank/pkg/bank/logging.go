package bank

import (
	"fmt"
	"github.com/rs/zerolog"
	"time"

	"github.com/CodingSquire/bank/pkg/api"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	svc    Bank
	logger *api.Logger
}

func (s *loggingMiddleware) AddToBell(request api.AddToBellRequest) (response api.AddToBellResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(err).
			Str("method", "AddToBell").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "AddToBell").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.AddToBell(request)
	return response, err
}

func (s *loggingMiddleware) DeductFromBell(request api.DeductFromBellRequest) (response api.DeductFromBellResponse, err error) {
	defer func(begin time.Time) {
		fmt.Println("___________________________")
		fmt.Println("method", "DeductFromBell")
		fmt.Println("timestamp", time.Now())
		fmt.Println("response", response)
		fmt.Println("elapsed", time.Since(begin))
		fmt.Println("err", err)
		fmt.Println("End.")
		fmt.Println("___________________________")
	}(time.Now())
	fmt.Println("___________________________")
	fmt.Println("method", "DeductFromBell")
	fmt.Println("timestamp", time.Now())
	fmt.Println("request", request)
	fmt.Println("Start.")
	fmt.Println("___________________________")
	response, err = s.svc.DeductFromBell(request)
	return response, err
}

func (s *loggingMiddleware) CreateAcc(request api.CreateAccRequest) (response api.CreateAccResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(err).
			Str("method", "CreateAcc").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "CreateAcc").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.CreateAcc(request)
	return response, err
}

func (s *loggingMiddleware) GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(err).
			Str("method", "GetBalance").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "GetBalance").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.GetBalance(request)
	return response, err
}

func (s *loggingMiddleware) wrap(err error) *zerolog.Event {
	lvl := s.logger.Info()
	if err != nil {
		lvl = s.logger.Error()
	}
	return lvl
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(svc Bank, logger *api.Logger) Bank {
	return &loggingMiddleware{
		svc:    svc,
		logger: logger,
	}
}
