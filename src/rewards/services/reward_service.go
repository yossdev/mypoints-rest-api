package services

import "github.com/yossdev/mypoints-rest-api/src/rewards/entities"

type rewardService struct {
	rewardRepository entities.PsqlRepository
}

func NewRewardService(p entities.PsqlRepository) entities.Service {
	return &rewardService{
		rewardRepository: p,
	}
}

func (s *rewardService) CreateReward() error {
	return nil
}
