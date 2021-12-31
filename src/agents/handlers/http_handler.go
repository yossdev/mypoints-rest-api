package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/agents/dto"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

// AgentHandlers contains method used for the handler
type AgentHandlers interface {
	SignIn(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	GetAgent(c *fiber.Ctx) error
	UpdateAgent(c *fiber.Ctx) error
	UpdateAvatar(c *fiber.Ctx) error
}

type agentHandler struct {
	agentService entities.Service
}

func NewHttpHandler(s entities.Service) *agentHandler {
	return &agentHandler{
		agentService: s,
	}
}

// SignIn post handler.
// @Description check agent by checking given email and password.
// @Summary check agent by given email return jwt token if successfully signIn
// @Tags Agent
// @Scheme https
// @Accept json
// @Produce json
// @Param signIn body dto.SignInReq true "body request"
// @Success 200 {object} auth.Token
// @Router /login [post]
func (h *agentHandler) SignIn(c *fiber.Ctx) error {
	payload := new(dto.SignInReq)
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

	res, err := h.agentService.SignIn(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, web.BadCredential, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Welcome, res)
}

// SignUp post handler.
// @Description create agent account by admins.
// @Summary admins can create agent account with this api
// @Tags Agent
// @Scheme https
// @Accept json
// @Produce json
// @Param signUp body dto.SignUpReq true "body request"
// @Success 201 {object} dto.AccountCreated
// @Router /:id/agent [post]
func (h *agentHandler) SignUp(c *fiber.Ctx) error {
	payload := new(dto.SignUpReq)
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

	res, err := h.agentService.SignUp(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusConflict, web.DuplicateData, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.AccountCreated, dto.FromDomainAC(res))
}

// GetAgent get handler.
// @Description Get agent data by id.
// @Summary get agent data
// @Tags Agent
// @Accept json
// @Produce json
// @Success 200 {object} dto.Profile
// @Router /profile/:id [get]
func (h *agentHandler) GetAgent(c *fiber.Ctx) error {
	id := c.Params("id")
	agent, err := h.agentService.GetAgent(uuid.MustParse(id))
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusForbidden, web.Forbidden, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainProfile(agent))
}

// UpdateAgent put handler.
// @Description update agent data by id.
// @Summary update agent data
// @Tags Agent
// @Accept json
// @Produce json
// @Param updateAccount body dto.UpdateAccount true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Router /profile/:id [put]
func (h *agentHandler) UpdateAgent(c *fiber.Ctx) error {
	payload := new(dto.UpdateAccount)
	id := c.Params("id")

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

	res, err := h.agentService.UpdateAgent(uuid.MustParse(id), payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusInternalServerError, web.InternalServerErr, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}

// UpdateAvatar put handler.
// @Description update agent photo profile by id.
// @Summary update agent photo profile
// @Tags Agent
// @Accept json
// @Produce json
// @Param updateAvatar body dto.UpdateAvatar true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Router /profile/avatar/:id [put]
func (h *agentHandler) UpdateAvatar(c *fiber.Ctx) error {
	payload := new(dto.UpdateAvatar)
	id := c.Params("id")

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

	res, err := h.agentService.UpdateAvatar(uuid.MustParse(id), payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusInternalServerError, web.InternalServerErr, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}
