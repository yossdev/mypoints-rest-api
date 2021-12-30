package repositories

import (
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

type transactionPsqlRepository struct {
	DB db.PsqlDB
}

func NewTransactionPsqlRepository(p db.PsqlDB) entities.PsqlRepository {
	return &transactionPsqlRepository{
		DB: p,
	}
}

func (p *transactionPsqlRepository) Create() error {
	return nil
}
