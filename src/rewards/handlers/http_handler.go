package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/rewards/dto"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
)

// RewardHandlers contains method used for the handler
type RewardHandlers interface {
	CreateReward(c *fiber.Ctx) error
	UpdateReward(c *fiber.Ctx) error
	DeleteReward(c *fiber.Ctx) error
}

type rewardHandlers struct {
	RewardService entities.Service
}

func NewHttpHandler(s entities.Service) RewardHandlers {
	return &rewardHandlers{
		RewardService: s,
	}
}

// CreateReward post handler.
// @Description create reward by admins.
// @Summary admins can create reward
// @Tags Reward
// @Scheme https
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param newReward body dto.NewReward true "body request"
// @Success 201 {object} dto.RewardRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 424 {object} web.ErrorResp
// @Router /reward/{id} [post]
// @Security ApiKeyAuth
func (h *rewardHandlers) CreateReward(c *fiber.Ctx) error {
	payload := new(dto.NewReward)
	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	// Create a new validator.
	validate := helpers.NewValidator()
	// Validate fields from payload.
	if err := validate.Struct(payload); err != nil {
		// Return, if some fields are not valid.
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	res, err := h.RewardService.CreateReward(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusFailedDependency, web.Failed, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.RewardCreated, dto.FromDomainRA(res))
}

// UpdateReward put handler.
// @Description update reward data by id.
// @Summary update reward data
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param rewardId path int true "ID of Reward"
// @Param updateReward body dto.UpdateReward true "body request"
// @Success 200 {object} dto.RewardRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 422 {object} web.ErrorResp
// @Router /reward/{id}/{rewardId} [put]
// @Security ApiKeyAuth
func (h *rewardHandlers) UpdateReward(c *fiber.Ctx) error {
	params := c.Params("rewardId")
	rewardId, convErr := helpers.StringToUint32(params)
	if convErr != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnprocessableEntity, web.CannotProcess, convErr)
	}

	payload := new(dto.UpdateReward)
	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	// Create a new validator.
	validate := helpers.NewValidator()
	// Validate fields from payload.
	if err := validate.Struct(payload); err != nil {
		// Return, if some fields are not valid.
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	res, err := h.RewardService.UpdateReward(rewardId, payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, fiber.ErrBadRequest)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainRA(res))
}

// DeleteReward delete handler.
// @Description soft delete reward data by id.
// @Summary soft delete reward data
// @Tags Reward
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param rewardId path int true "ID of Reward"
// @Success 200 {object} dto.RewardRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 422 {object} web.ErrorResp
// @Router /reward/{id}/{rewardId} [delete]
// @Security ApiKeyAuth
func (h *rewardHandlers) DeleteReward(c *fiber.Ctx) error {
	params := c.Params("rewardId")
	rewardId, convErr := helpers.StringToUint32(params)
	if convErr != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnprocessableEntity, web.CannotProcess, convErr)
	}

	res, err := h.RewardService.DeleteReward(rewardId)
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, fiber.ErrBadRequest)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainRA(res))
}
