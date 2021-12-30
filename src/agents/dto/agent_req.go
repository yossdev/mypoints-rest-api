package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

type SignInReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *SignInReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

type SignUpReq struct {
	AdminID  uuid.UUID `json:"admin_id" validate:"required,uuid"`
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required"`
	Img      string    `json:"img"`
	Active   bool      `json:"active"`
}

func (req *SignUpReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		AdminID:  req.AdminID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Img:      req.Img,
		Active:   req.Active,
	}
}

type UpdateAccount struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (req *UpdateAccount) ToDomain() *entities.Domain {
	return &entities.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type UpdateAvatar struct {
	Img string `json:"img" validate:"required"`
}

func (req *UpdateAvatar) ToDomain() *entities.Domain {
	return &entities.Domain{
		Img: req.Img,
	}
}
