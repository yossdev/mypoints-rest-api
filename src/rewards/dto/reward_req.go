package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

type NewReward struct {
	AdminID  uuid.UUID `json:"admin_id" validate:"required,uuid"`
	Title    string    `json:"title" validate:"required"`
	Value    uint64    `json:"value" validate:"required,numeric,gt=0"`
	Points   uint32    `json:"points" validate:"required,numeric,gt=0"`
	Img      string    `json:"img" validate:"omitempty,url"`
	Category string    `json:"category" validate:"required"`
}

func (req *NewReward) ToDomain() entities.Domain {
	return entities.Domain{
		AdminID:  req.AdminID,
		Title:    req.Title,
		Value:    req.Value,
		Points:   req.Points,
		Img:      req.Img,
		Category: req.Category,
	}
}

type UpdateReward struct {
	Title    string `json:"title" validate:"required"`
	Value    uint64 `json:"value" validate:"required,numeric,gt=0"`
	Points   uint32 `json:"points" validate:"required,numeric,gt=0"`
	Img      string `json:"img" validate:"omitempty,url"`
	Category string `json:"category" validate:"required"`
}

func (req *UpdateReward) ToDomain() entities.Domain {
	return entities.Domain{
		Title:    req.Title,
		Value:    req.Value,
		Points:   req.Points,
		Img:      req.Img,
		Category: req.Category,
	}
}
