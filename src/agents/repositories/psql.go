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
	agent := Agent{}

	res := p.DB.DB().Where("email = ?", email).First(&agent)
	if res.Error != nil {
		return agent.toDomain(), res.Error
	}

	return agent.toDomain(), nil
}

func (p *agentPsqlRepository) GetAgent(id uuid.UUID) (*entities.Domain, error) {
	agent := Agent{}
	if err := p.DB.DB().First(&agent, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return agent.toDomain(), nil
}

func (p *agentPsqlRepository) CreateAgent(payload *entities.Domain, adminID uuid.UUID) (int64, error) {
	agent := Agent{
		AdminID:  adminID,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Img:      payload.Img,
		Status:   payload.Status,
	}
	res := p.DB.DB().Create(&agent)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

//func (p *agentPsqlRepository) UpdateAgent(payload *entities.Domain) error {
//	return nil
//}
