package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

type ClaimsReq struct {
	AgentID   uuid.UUID `json:"agent_id"  validate:"required,uuid"`
	ProductID uint32    `json:"product_id" validate:"required,numeric,gt=0"`
	Title     string    `json:"title" validate:"required"`
	Points    uint32    `json:"points" validate:"required,numeric,gt=0"`
	NotaImg   string    `json:"nota_img" validate:"required,url"`
}

func (req *ClaimsReq) ToDomain() entities.Domain {
	return entities.Domain{
		AgentID:   req.AgentID,
		ProductID: req.ProductID,
		Title:     req.Title,
		Points:    req.Points,
		NotaImg:   req.NotaImg,
	}
}

type RedeemReq struct {
	AgentID    uuid.UUID `json:"agent_id"  validate:"required,uuid"`
	RewardID   uint32    `json:"reward_id" validate:"required,numeric,gt=0"`
	Title      string    `json:"title" validate:"required"`
	Points     uint32    `json:"points" validate:"required,numeric,gt=0"`
	RedeemDesc string    `json:"redeem_desc" validate:"required"`
}

func (req *RedeemReq) ToDomain() entities.Domain {
	return entities.Domain{
		AgentID:    req.AgentID,
		RewardID:   req.RewardID,
		Title:      req.Title,
		Points:     req.Points,
		RedeemDesc: req.RedeemDesc,
	}
}

type UpdateClaimsReq struct {
	ID     uuid.UUID `json:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}

type InvoiceCallback struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	MerchantName string `json:"merchant_name"`
	Amount       int    `json:"amount"`
}

func (req *InvoiceCallback) ToDomain() entities.InvoiceCallback {
	return entities.InvoiceCallback{
		ID:           req.ID,
		Status:       req.Status,
		MerchantName: req.MerchantName,
		Amount:       req.Amount,
	}
}
