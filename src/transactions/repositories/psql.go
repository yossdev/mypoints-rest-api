package repositories

import (
	"github.com/google/uuid"
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

func (p *transactionPsqlRepository) CreateClaims(payload entities.Domain) (int64, error) {
	claims := Transaction{}
	createClaims(payload, &claims)

	res := p.DB.DB().Omit("RewardID", "RedeemInvoice").Create(&claims)
	return res.RowsAffected, res.Error
}

func (p *transactionPsqlRepository) UpdateClaimsStatus(id uuid.UUID, status string) (int64, error) {
	claims := Transaction{Status: status}

	res := p.DB.DB().Model(&claims).Where("id = ?", id).Updates(claims)
	return res.RowsAffected, res.Error
}
