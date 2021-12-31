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

func (p *rewardPsqlRepository) Create(payload entities.Domain) (int64, error) {
	reward := Reward{}
	createReward(payload, &reward)

	res := p.DB.DB().Create(&reward)
	return res.RowsAffected, res.Error
}

func (p *rewardPsqlRepository) Update(payload entities.Domain) (int64, error) {
	reward := Reward{}
	updateReward(payload, &reward)

	res := p.DB.DB().Model(&reward).Where("id = ?", payload.ID).Updates(reward)
	return res.RowsAffected, res.Error
}

func (p *rewardPsqlRepository) Delete(productId uint32) (int64, error) {
	res := p.DB.DB().Delete(&Reward{}, productId)
	return res.RowsAffected, res.Error
}
