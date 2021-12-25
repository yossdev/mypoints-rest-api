package dto

import "github.com/yossdev/mypoints-rest-api/src/admins/entities"

type SignInReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req *SignInReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Img      string `json:"img"`
}

func (req *SignUpReq) ToDomain() *entities.Domain {
	return &entities.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Img:      req.Img,
	}
}
