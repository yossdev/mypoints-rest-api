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
}

func (req *ClaimsReq) ToDomain() entities.Domain {
	return entities.Domain{
		AgentID:   req.AgentID,
		ProductID: req.ProductID,
		Title:     req.Title,
		Points:    req.Points,
	}
}

type RedeemReq struct {
	AgentID  uuid.UUID `json:"agent_id"  validate:"required,uuid"`
	RewardID uint32    `json:"reward_id" validate:"required,numeric,gt=0"`
	Title    string    `json:"title" validate:"required"`
	Points   uint32    `json:"points" validate:"required,numeric,gt=0"`
}

func (req *RedeemReq) ToDomain() entities.Domain {
	return entities.Domain{
		AgentID:  req.AgentID,
		RewardID: req.RewardID,
		Title:    req.Title,
		Points:   req.Points,
	}
}

type UpdateClaimsReq struct {
	ID     uuid.UUID `json:"id" validate:"required,uuid"`
	Status string    `json:"status" validate:"required"`
}
