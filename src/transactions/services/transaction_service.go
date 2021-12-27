package services

import "github.com/yossdev/mypoints-rest-api/src/transactions/entities"

type transactionService struct {
	transactionPsqlRepository entities.PsqlRepository
}

func NewTransactionService(p entities.PsqlRepository) entities.Service {
	return &transactionService{
		transactionPsqlRepository: p,
	}
}

func (s *transactionService) GetTransactions() error {
	return nil
}

func (s *transactionService) GetTransactionDetail() error {
	return nil
}

func (s *transactionService) GetTransactionsAdmin() error {
	return nil
}