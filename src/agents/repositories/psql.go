package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"gorm.io/gorm"
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
		return entities.Domain{}, res.Error
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
		return entities.Domain{}, err
	}

	return agent.ToDomain(), nil
}

func (p *agentPsqlRepository) CreateAgent(payload *entities.Domain) (int64, error) {
	agent := Agent{}
	createAccount(payload, &agent)

	res := p.DB.DB().Create(&agent)
	return res.RowsAffected, res.Error
}

func (p *agentPsqlRepository) UpdateAgent(payload entities.Domain) (int64, error) {
	agent := Agent{}
	updateAccount(payload, &agent)

	res := p.DB.DB().Model(&agent).Where("id = ?", payload.ID).Updates(agent)
	return res.RowsAffected, res.Error
}

func (p *agentPsqlRepository) UpdateAvatar(payload entities.Domain) (int64, error) {
	res := p.DB.DB().Model(&Agent{}).Where("id = ?", payload.ID).Update("img", payload.Img)
	return res.RowsAffected, res.Error
}

func (p *agentPsqlRepository) UpdatePoints(id uuid.UUID, points int32) (int64, error) {
	agent := Agent{}

	res := p.DB.DB().Model(&agent).Where("id = ?", id).Update("points", gorm.Expr("points + ?", points))
	return res.RowsAffected, res.Error
}

func (p *agentPsqlRepository) UpdateAgentByAdmin(payload entities.Domain) (int64, error) {
	ignore := "password"
	if payload.Password != "" {
		ignore = ""
	}

	res := p.DB.DB().Model(&Agent{}).Omit(ignore).Where("id = ?", payload.ID).Updates(map[string]interface{}{"password": payload.Password, "active": payload.Active})
	return res.RowsAffected, res.Error
}
