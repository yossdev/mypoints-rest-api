package dto

import "github.com/yossdev/mypoints-rest-api/src/admins/entities"

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
