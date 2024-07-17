package model

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrNotEnoughMoney  = errors.New("not enough money")
)
