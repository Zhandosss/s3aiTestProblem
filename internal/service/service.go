package service

import "bankomat/internal/model"

type IRepository interface {
	CreateAccount() (string, error)
	Deposit(operation *model.Operation) error
	Withdraw(operation *model.Operation) error
	GetBalance(id string) (float64, error)
}

type Service struct {
	repository IRepository
}

func New(repository IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateAccount(idChan chan<- string, errorChan chan<- error) {
	id, err := s.repository.CreateAccount()
	if err != nil {
		errorChan <- err
		return
	}

	idChan <- id
	return
}

func (s *Service) Deposit(operationChan <-chan *model.Operation, errorChan chan<- error) {
	operation := <-operationChan

	err := s.repository.Deposit(operation)
	if err != nil {
		errorChan <- err
		return
	}

	errorChan <- nil
}

func (s *Service) Withdraw(operationChan <-chan *model.Operation, errorChan chan<- error) {
	operation := <-operationChan

	err := s.repository.Withdraw(operation)
	if err != nil {
		errorChan <- err
		return
	}

	errorChan <- nil
}

func (s *Service) GetBalance(idChan <-chan string, balanceChan chan<- *model.BalanceFromService) {
	id := <-idChan

	balance, err := s.repository.GetBalance(id)
	if err != nil {
		balanceChan <- &model.BalanceFromService{
			Balance: 0,
			Err:     err,
		}
		return
	}

	balanceChan <- &model.BalanceFromService{
		Balance: balance,
		Err:     nil,
	}
}
