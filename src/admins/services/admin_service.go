package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
)

type adminService struct {
	adminPsqlRepository entities.PsqlRepository
}

func NewAdminService(p entities.PsqlRepository) entities.Service {
	return &adminService{
		adminPsqlRepository: p,
	}
}

func (s *adminService) SignIn(payload *entities.Domain) (auth.Token, error) {
	admin, err := s.adminPsqlRepository.SignInWithEmail(payload.Email)
	if err != nil {
		return auth.Token{}, err
	}

	if err := helpers.ValidateHash(admin.Password, payload.Password); err != nil {
		return auth.Token{}, err
	}

	token := auth.Sign(admin.ID, jwt.MapClaims{
		"sub": admin.ID,
		"https://hasura.io/jwt/claims": fiber.Map{
			"x-hasura-default-role": "admins",
			// do some custom logic to decide allowed roles
			"x-hasura-allowed-roles": []string{"admins"},
			"x-hasura-admins-id":     admin.ID,
		},
	})

	return token, nil
}
