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

func (p *transactionPsqlRepository) GetTransaction(id string) (entities.Domain, error) {
	transaction := Transaction{}

	var err error
	param, e := uuid.Parse(id)
	if e != nil {
		err = p.DB.DB().Where("redeem_invoice_id = ?", id).First(&transaction).Error
	} else {
		err = p.DB.DB().Where("id = ?", param).First(&transaction).Error
	}
	if err != nil {
		return entities.Domain{}, err
	}

	return transaction.ToDomain(), nil
}

func (p *transactionPsqlRepository) CreateClaims(payload entities.Domain) (int64, error) {
	claims := Transaction{}
	createClaims(payload, &claims)

	res := p.DB.DB().Omit("RewardID", "RedeemInvoiceID", "RedeemInvoiceURL", "RedeemDesc").Create(&claims)
	return res.RowsAffected, res.Error
}

func (p *transactionPsqlRepository) UpdateClaimsStatus(id uuid.UUID, status string) (int64, error) {
	claims := Transaction{Status: status}

	res := p.DB.DB().Model(&claims).Where("id = ?", id).Updates(claims)
	return res.RowsAffected, res.Error
}

func (p *transactionPsqlRepository) CreateRedeem(payload entities.Domain) (int64, error) {
	redeem := Transaction{}
	createRedeem(payload, &redeem)

	res := p.DB.DB().Omit("ProductID", "NotaImg").Create(&redeem)
	return res.RowsAffected, res.Error
}

func (p *transactionPsqlRepository) UpdateRedeemStatus(id, status string) error {
	redeem := Transaction{Status: status}
	err := p.DB.DB().Model(&redeem).Where("redeem_invoice_id = ?", id).Updates(redeem)

	return err.Error
}
