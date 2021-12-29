package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/rewards/dto"
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

// CreateReward post handler.
// @Description create reward by admins.
// @Summary admins can create reward
// @Tags Admin
// @Scheme https
// @Accept json
// @Produce json
// @Param newReward body dto.NewReward true "body request"
// @Success 201 {object} dto.RowsAffected
// @Router /reward/:id [post]
func (h *rewardHandler) CreateReward(c *fiber.Ctx) error {
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

	res, err := h.rewardService.CreateReward(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusFailedDependency, web.Failed, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.RewardCreated, dto.RowsAffected{RowsAffected: res})
}

// UpdateReward put handler.
// @Description update reward data by id.
// @Summary update reward data
// @Tags Reward
// @Accept json
// @Produce json
// @Param updateReward body dto.UpdateReward true "body request"
// @Success 200 {object} dto.RowsAffected
// @Router /reward/:id/:rewardId [put]
func (h *rewardHandler) UpdateReward(c *fiber.Ctx) error {
	payload := new(dto.UpdateReward)
	rewardId := c.Params("rewardId")

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

	res, err := h.rewardService.UpdateReward(rewardId, payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnprocessableEntity, fiber.ErrUnprocessableEntity, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.RowsAffected{RowsAffected: res})
}

// DeleteReward delete handler.
// @Description soft delete reward data by id.
// @Summary soft delete reward data
// @Tags Reward
// @Accept json
// @Produce json
// @Param deleteReward body dto.DeleteReward true "body request"
// @Success 200 {object} dto.RowsAffected
// @Router /reward/:id [delete]
func (h *rewardHandler) DeleteReward(c *fiber.Ctx) error {
	payload := new(dto.DeleteReward)
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

	res, err := h.rewardService.DeleteReward(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusNotFound, web.IDNotFound, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.RowsAffected{RowsAffected: res})
}
