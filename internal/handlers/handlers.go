package handlers

import (
	"bankomat/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type IService interface {
	CreateAccount(idChan chan<- string, errorChan chan<- error)
	Deposit(operationChan <-chan *model.Operation, errorChan chan<- error)
	Withdraw(operationChan <-chan *model.Operation, errorChan chan<- error)
	GetBalance(idChan <-chan string, balanceChan chan<- *model.BalanceFromService)
}

type Handlers struct {
	service IService
}

func New(e *echo.Echo, service IService) {

	h := &Handlers{
		service: service,
	}

	e.Use(middleware.RequestID())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogRequestID: true,
		LogRemoteIP:  true,
		LogError:     true,
		LogLatency:   true,
		LogMethod:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("latentcy", v.Latency.String()).
				Str("requestID", v.RequestID).
				Err(v.Error).
				Str("remoteIP", v.RemoteIP).
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("method", v.Method).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.POST("/accounts", h.CreateAccount)
	e.POST("/accounts/:id/deposit", h.Deposit)
	e.POST("/accounts/:id/withdraw", h.Withdraw)
	e.GET("/accounts/:id/balance", h.GetBalance)
}
