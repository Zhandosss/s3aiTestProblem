package model

import "sync"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	Balance float64
	mx      sync.Mutex
}

func (a *Account) Deposit(amount float64) error {
	a.mx.Lock()
	defer a.mx.Unlock()

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mx.Lock()
	defer a.mx.Unlock()

	if a.Balance < amount {
		return ErrNotEnoughMoney
	}

	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mx.Lock()
	defer a.mx.Unlock()

	return a.Balance
}

// Operation represents a money operation(deposit/withdraw) for a user
type Operation struct {
	Amount float64 `json:"amount"`
	UserID string  `json:"user_id"`
}

type MoneyTransfer struct {
	Amount float64 `json:"amount"`
}

type BalanceFromService struct {
	Balance float64 `json:"balance"`
	Err     error   `json:"error"`
}
