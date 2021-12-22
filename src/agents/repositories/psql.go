package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

type agentPsqlRepository struct {
	DB db.PsqlDB
}

func NewAgentPsqlRepository(p db.PsqlDB) entities.PsqlRepository {
	return &agentPsqlRepository{
		DB: p,
	}
}

func (p *agentPsqlRepository) SignInWithEmail(email string) (*entities.Domain, error) {
	agent := Agents{}

	res := p.DB.DB().Where("email = ?", email).First(&agent)
	if res.Error != nil {
		return agent.toDomain(), res.Error
	}

	return agent.toDomain(), nil
}

func (p *agentPsqlRepository) GetAgent(id uuid.UUID) (*entities.Domain, error) {
	agent := Agents{}
	//if err := u.DB.DB().First(&agent, "id = ?", id).Error; err != nil {
	//	return nil, err
	//}

	return agent.toDomain(), nil
}

func (p *agentPsqlRepository) UpdateAgent(payload *entities.Domain) error {
	return nil
}
