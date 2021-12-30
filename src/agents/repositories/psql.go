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

func (p *agentPsqlRepository) SignInWithEmail(email string) (entities.Domain, error) {
	agent := Agent{}

	res := p.DB.DB().Where("email = ?", email).First(&agent)
	if res.Error != nil {
		return agent.ToDomain(), res.Error
	}

	return agent.ToDomain(), nil
}

func (p *agentPsqlRepository) GetAgent(id uuid.UUID) (entities.Domain, error) {
	agent := Agent{}

	// Preload all relations
	//if err := p.DB.DB().Preload(clause.Associations).First(&agent, "id = ?", id).Error; err != nil {
	//	return agent.ToDomain(), err
	//}

	if err := p.DB.DB().First(&agent, "id = ?", id).Error; err != nil {
		return agent.ToDomain(), err
	}

	return agent.ToDomain(), nil
}

func (p *agentPsqlRepository) CreateAgent(payload *entities.Domain) (int64, error) {
	agent := Agent{
		AdminID:  payload.AdminID,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Img:      payload.Img,
		Active:   payload.Active,
	}
	res := p.DB.DB().Create(&agent)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *agentPsqlRepository) UpdateAgent(payload *entities.Domain) (int64, error) {
	agent := Agent{}
	p.DB.DB().First(&agent, "id = ?", payload.ID)

	updateAccount(payload, &agent)

	res := p.DB.DB().Save(&agent)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *agentPsqlRepository) UpdateAvatar(payload *entities.Domain) (int64, error) {
	agent := Agent{}
	p.DB.DB().First(&agent, "id = ?", payload.ID)

	agent.Img = payload.Img

	res := p.DB.DB().Save(&agent)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
