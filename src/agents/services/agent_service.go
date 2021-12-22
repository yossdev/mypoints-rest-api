package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
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

func (s *agentService) SignIn(payload *entities.Domain) (interface{}, error) {
	agent, err := s.agentPsqlRepository.SignInWithEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := helpers.ValidateHash(agent.Password, payload.Password); err != nil {
		return nil, err
	}

	token := auth.Sign(agent.ID, jwt.MapClaims{
		"id": agent.ID,
	})

	return token, nil
}

func (s *agentService) GetAgent(id uuid.UUID) (*entities.Domain, error) {
	agent, err := s.agentPsqlRepository.GetAgent(id)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (s *agentService) UpdateAgent(id uuid.UUID, payload *entities.Domain) error {
	return nil
}
