package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
)

type NewProduct struct {
	AdminID uuid.UUID `json:"admin_id" validate:"required,uuid"`
	Title   string    `json:"title" validate:"required"`
	Points  uint32    `json:"points" validate:"required,numeric,gt=0"`
	Img     string    `json:"img" validate:"omitempty,url"`
}

func (req *NewProduct) ToDomain() entities.Domain {
	return entities.Domain{
		AdminID: req.AdminID,
		Title:   req.Title,
		Points:  req.Points,
		Img:     req.Img,
	}
}

type UpdateProduct struct {
	Title  string `json:"title" validate:"required"`
	Points uint32 `json:"points" validate:"required,numeric,gt=0"`
	Img    string `json:"img" validate:"omitempty,url"`
}

func (req *UpdateProduct) ToDomain() entities.Domain {
	return entities.Domain{
		Title:  req.Title,
		Points: req.Points,
		Img:    req.Img,
	}
}
