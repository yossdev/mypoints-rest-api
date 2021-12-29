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

func (p *rewardPsqlRepository) Create() error {
	return nil
}
