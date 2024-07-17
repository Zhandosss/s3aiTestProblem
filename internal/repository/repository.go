package repository

import (
	"bankomat/internal/model"
	"fmt"
)

type Repository struct {
	db    map[string]*model.Account
	count int
}

func New() *Repository {
	return &Repository{
		db:    make(map[string]*model.Account),
		count: 1,
	}
}

func (r *Repository) CreateAccount() (string, error) {
	id := fmt.Sprintf("%d", r.count)

	r.db[id] = &model.Account{
		Balance: 0,
	}

	r.count++

	return id, nil
}

func (r *Repository) Deposit(operation *model.Operation) error {
	if _, ok := r.db[operation.UserID]; !ok {
		return model.ErrAccountNotFound
	}

	err := r.db[operation.UserID].Deposit(operation.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Withdraw(operation *model.Operation) error {
	if _, ok := r.db[operation.UserID]; !ok {
		return model.ErrAccountNotFound
	}

	err := r.db[operation.UserID].Withdraw(operation.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetBalance(id string) (float64, error) {
	if _, ok := r.db[id]; !ok {
		return 0, model.ErrAccountNotFound
	}

	return r.db[id].GetBalance(), nil
}
