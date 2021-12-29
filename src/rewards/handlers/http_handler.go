package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

type rewardHandler struct {
	rewardService entities.Service
}

func NewHttpHandler(s entities.Service) *rewardHandler {
	return &rewardHandler{
		rewardService: s,
	}
}

func (h *rewardHandler) CreateReward(c *fiber.Ctx) error {
	return nil
}
