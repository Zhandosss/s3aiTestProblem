package handlers

import (
	"bankomat/internal/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) CreateAccount(c echo.Context) error {
	idChan := make(chan string)
	errorChan := make(chan error)

	go h.service.CreateAccount(idChan, errorChan)

	select {
	case id := <-idChan:
		log.Info().Msgf("Account created with id: %s", id)
		return c.JSON(http.StatusCreated, map[string]string{"id": id})
	case err := <-errorChan:
		log.Error().Err(err).Msg("error creating account")

		if errors.Is(err, model.ErrAccountNotFound) {
			return c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "account already exists"})
		}

		return c.JSON(http.StatusInternalServerError, &ErrorResponse{Message: "error creating account"})
	}
}
