package repositories

import (
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

type rewardPsqlRepository struct {
	DB db.PsqlDB
}

func NewRewardPsqlRepository(p db.PsqlDB) entities.PsqlRepository {
	return &rewardPsqlRepository{
		DB: p,
	}
}

func (p *rewardPsqlRepository) Create(payload *entities.Domain) (int64, error) {
	reward := Reward{
		AdminID: payload.AdminID,
		Title:   payload.Title,
		Value:   payload.Value,
		Points:  payload.Points,
		Img:     payload.Img,
	}

	res := p.DB.DB().Create(&reward)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *rewardPsqlRepository) Update(payload *entities.Domain) (int64, error) {
	reward := Reward{}
	if err := p.DB.DB().First(&reward, "id = ?", payload.ID); err.Error != nil {
		return 0, err.Error
	}

	updateReward(payload, &reward)

	res := p.DB.DB().Save(&reward)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *rewardPsqlRepository) Delete(id uint32) (int64, error) {
	reward := Reward{}
	if err := p.DB.DB().First(&reward, "id = ?", id).Error; err != nil {
		return 0, err
	}

	res := p.DB.DB().Delete(&reward)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
