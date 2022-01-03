package services

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

type transactionService struct {
	transactionPsqlRepository entities.PsqlRepository
}

func NewTransactionService(p entities.PsqlRepository) entities.Service {
	return &transactionService{
		transactionPsqlRepository: p,
	}
}

func (s *transactionService) Claims(payload entities.Domain) (int64, error) {
	res, err := s.transactionPsqlRepository.CreateClaims(payload)
	return res, err
}

func (s *transactionService) ClaimsStatus(id uuid.UUID, status string) (int64, error) {
	res, err := s.transactionPsqlRepository.UpdateClaimsStatus(id, status)
	return res, err
}
