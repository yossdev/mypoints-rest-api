package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

type agentService struct {
	agentPsqlRepository entities.PsqlRepository
}

func NewAgentService(p entities.PsqlRepository) entities.Service {
	return &agentService{
		agentPsqlRepository: p,
	}
}

func (s *agentService) SignIn(payload entities.Domain) (auth.Token, error) {
	agent, err := s.agentPsqlRepository.SignInWithEmail(payload.Email)
	if err != nil {
		return auth.Token{}, err
	}

	if !agent.Active {
		return auth.Token{}, web.AccountDisabled
	}

	if err := helpers.ValidateHash(agent.Password, payload.Password); err != nil {
		return auth.Token{}, err
	}

	token := auth.Sign(jwt.MapClaims{
		"sub": agent.ID,
		"https://hasura.io/jwt/claims": fiber.Map{
			"x-hasura-default-role": "agent",
			// do some custom logic to decide allowed roles
			"x-hasura-allowed-roles": []string{"agent"},
			"x-hasura-agent-id":      agent.ID,
		},
		"role": "agent",
	})

	return token, nil
}

func (s *agentService) GetAgent(id uuid.UUID) (entities.Domain, error) {
	agent, err := s.agentPsqlRepository.GetAgent(id)
	return agent, err
}

func (s *agentService) SignUp(payload *entities.Domain) (int64, error) {
	payload.Password, _ = helpers.Hash(payload.Password)
	res, err := s.agentPsqlRepository.CreateAgent(payload)

	return res, err
}

func (s *agentService) UpdateAgent(id uuid.UUID, payload entities.Domain) (int64, error) {
	payload.ID = id
	if payload.Password != "" {
		payload.Password, _ = helpers.Hash(payload.Password)
	}

	res, err := s.agentPsqlRepository.UpdateAgent(payload)

	return res, err
}

func (s *agentService) UpdateAvatar(id uuid.UUID, payload entities.Domain) (int64, error) {
	payload.ID = id
	res, err := s.agentPsqlRepository.UpdateAvatar(payload)

	return res, err
}

func (s *agentService) UpdateAgentByAdmin(payload entities.Domain) (int64, error) {
	if payload.Password != "" {
		payload.Password, _ = helpers.Hash(payload.Password)
	}

	return s.agentPsqlRepository.UpdateAgentByAdmin(payload)
}
