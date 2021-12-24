package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *SignInReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

type SignUpReq struct {
	AdminID  uuid.UUID `json:"admin_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Img      string    `json:"img"`
}

func (req *SignUpReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		AdminID:  req.AdminID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Img:      req.Img,
	}
}
