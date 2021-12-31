package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func (s *adminService) SignIn(payload entities.Domain) (auth.Token, error) {
	admin, err := s.adminPsqlRepository.SignInWithEmail(payload.Email)
	if err != nil {
		return auth.Token{}, err
	}

	if err := helpers.ValidateHash(admin[0], payload.Password); err != nil {
		return auth.Token{}, err
	}

	token := auth.Sign(jwt.MapClaims{
		"sub": admin[1],
		"https://hasura.io/jwt/claims": fiber.Map{
			"x-hasura-default-role": "admins",
			// do some custom logic to decide allowed roles
			"x-hasura-allowed-roles": []string{"admins"},
			"x-hasura-admins-id":     admin[1],
		},
		"role": "admins",
	})

	return token, nil
}

func (s *adminService) SignUp(payload *entities.Domain) (int64, error) {
	payload.Password, _ = helpers.Hash(payload.Password)
	res, err := s.adminPsqlRepository.CreateAdmin(payload)

	return res, err
}

func (s *adminService) UpdateAdmin(id uuid.UUID, payload entities.Domain) (int64, error) {
	payload.ID = id
	if payload.Password != "" {
		payload.Password, _ = helpers.Hash(payload.Password)
	}

	res, err := s.adminPsqlRepository.UpdateAdmin(payload)

	return res, err
}

func (s *adminService) UpdateAvatar(id uuid.UUID, payload entities.Domain) (int64, error) {
	payload.ID = id
	res, err := s.adminPsqlRepository.UpdateAvatar(payload)

	return res, err
}
