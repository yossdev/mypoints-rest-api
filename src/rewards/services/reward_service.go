package services

import (
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

type rewardService struct {
	rewardPsqlRepository entities.PsqlRepository
}

func NewRewardService(p entities.PsqlRepository) entities.Service {
	return &rewardService{
		rewardPsqlRepository: p,
	}
}

func (s *rewardService) CreateReward(payload *entities.Domain) (int64, error) {
	res, err := s.rewardPsqlRepository.Create(payload)
	return res, err
}

func (s *rewardService) UpdateReward(rewardId uint32, payload *entities.Domain) (int64, error) {
	payload.ID = rewardId

	res, err := s.rewardPsqlRepository.Update(payload)
	return res, err
}

func (s *rewardService) DeleteReward(rewardId uint32) (int64, error) {
	res, err := s.rewardPsqlRepository.Delete(rewardId)
	return res, err
}
