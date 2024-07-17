package handlers

import (
	"bankomat/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) Deposit(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("invalid id")
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "invalid id"})
	}

	var moneyTransfer model.MoneyTransfer
	if err := c.Bind(&moneyTransfer); err != nil {
		log.Error().Err(err).Msgf("user %s. Invalid request body", id)
		return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "invalid request body. Please provide amount to deposit"})
	}
	log.Info().Msgf("user %s. Want  to deposit: %f", id, moneyTransfer.Amount)

	operationChan := make(chan *model.Operation)
	errorChan := make(chan error)

	go h.service.Deposit(operationChan, errorChan)

	operationChan <- &model.Operation{
		UserID: id,
		Amount: moneyTransfer.Amount,
	}

	err := <-errorChan
	if err != nil {
		log.Error().Err(err).Msgf("user %s. Error depositing money", id)

		if errors.Is(err, model.ErrAccountNotFound) {
			return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "account not found"})
		}

		return c.JSON(http.StatusInternalServerError, &ErrorResponse{Message: "error depositing money"})
	}

	log.Info().Msgf("user %s. Money deposited", id)
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}
