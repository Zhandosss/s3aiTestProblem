package handlers

import (
	"bankomat/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) GetBalance(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("invalid id")
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "invalid id"})
	}
	log.Info().Msgf("user %s. Want to get balance", id)

	idChan := make(chan string)
	balanceChan := make(chan *model.BalanceFromService)

	go h.service.GetBalance(idChan, balanceChan)

	idChan <- id
	balanceInfo := <-balanceChan

	if balanceInfo.Err != nil {
		log.Error().Err(balanceInfo.Err).Msgf("user %s. Error getting balance", id)

		if errors.Is(balanceInfo.Err, model.ErrAccountNotFound) {
			return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "account not found"})
		}

		return c.JSON(http.StatusInternalServerError, &ErrorResponse{Message: "error getting balance"})
	}

	log.Info().Msgf("user %s. Balance: %f", id, balanceInfo.Balance)
	return c.JSON(http.StatusOK, map[string]any{"amount": balanceInfo.Balance})
}
