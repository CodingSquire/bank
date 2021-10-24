package main

import (
	"fmt"
	"github.com/CodingSquire/bank/pkg/bank"
	"github.com/CodingSquire/bank/pkg/bank/httpserver"
	"github.com/CodingSquire/bank/pkg/logger"
	"github.com/valyala/fasthttp"
)

type configuration struct {
	Logger         logger.Config
	PortConference string `env:"PORT_CONFERENCE,default=8082"`
	Debug          bool   `env:"DEBUG,default=true"`
}

func main() {
	cfg := &configuration{
		Logger: logger.Config{
			Level:     "debug",
			Timestamp: true,
			Caller:    true,
			Pretty:    true,
		},
		Debug: true,
	}

	myBank := bank.NewBank()

	l := logger.NewLogger(&cfg.Logger)

	myBank = bank.NewLoggingMiddleware(myBank, l)
	myBankRouter := httpserver.NewPreparedServer(myBank)

	portMyBankRouter := ":9999"
	err := fasthttp.ListenAndServe(portMyBankRouter, myBankRouter.Handler)
	fmt.Println(err)

	fmt.Scanln()
}
