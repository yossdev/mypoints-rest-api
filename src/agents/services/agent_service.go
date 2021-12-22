package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"golang.org/x/crypto/bcrypt"
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

	if err := bcrypt.CompareHashAndPassword([]byte(agent.Password), []byte(payload.Password)); err != nil {
		return nil, err
	}

	return fiber.Map{"token": "berhasil login dan ini token"}, nil
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
