package services

import (
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	"strconv"
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

func (s *rewardService) UpdateReward(rewardId string, payload *entities.Domain) (int64, error) {
	ui64, err := strconv.ParseUint(rewardId, 10, 32)
	if err != nil {
		return 0, err
	}

	id := uint32(ui64)
	payload.ID = id

	res, err := s.rewardPsqlRepository.Update(payload)
	return res, err
}

func (s *rewardService) DeleteReward(payload *entities.Domain) (int64, error) {
	res, err := s.rewardPsqlRepository.Delete(payload.ID)
	return res, err
}
