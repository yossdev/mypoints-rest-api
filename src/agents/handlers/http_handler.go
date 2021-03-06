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
	UpdateAgentByAdmin(c *fiber.Ctx) error
}

type agentHandlers struct {
	AgentService entities.Service
}

func NewHttpHandler(s entities.Service) AgentHandlers {
	return &agentHandlers{
		AgentService: s,
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
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /login [post]
func (h *agentHandlers) SignIn(c *fiber.Ctx) error {
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

	res, err := h.AgentService.SignIn(payload.ToDomain())
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
// @Param adminId path string true "ID of Admin"
// @Param signUp body dto.SignUpReq true "body request"
// @Success 201 {object} dto.AccountCreated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 409 {object} web.ErrorResp
// @Router /{adminId}/agent [post]
// @Security ApiKeyAuth
func (h *agentHandlers) SignUp(c *fiber.Ctx) error {
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

	res, err := h.AgentService.SignUp(payload.ToDomain())
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
// @Param id path string true "ID of Agent"
// @Success 200 {object} dto.Profile
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 403 {object} web.ErrorResp
// @Router /profile/{id} [get]
// @Security ApiKeyAuth
func (h *agentHandlers) GetAgent(c *fiber.Ctx) error {
	id := c.Params("id")
	agent, err := h.AgentService.GetAgent(uuid.MustParse(id))
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
// @Param id path string true "ID of Agent"
// @Param updateAccount body dto.UpdateAccount true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /profile/{id} [put]
// @Security ApiKeyAuth
func (h *agentHandlers) UpdateAgent(c *fiber.Ctx) error {
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

	res, err := h.AgentService.UpdateAgent(uuid.MustParse(id), payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}

// UpdateAvatar put handler.
// @Description update agent photo profile by id.
// @Summary update agent photo profile
// @Tags Agent
// @Accept json
// @Produce json
// @Param id path string true "ID of Agent"
// @Param updateAvatar body dto.UpdateAvatar true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /profile/avatar/{id} [put]
// @Security ApiKeyAuth
func (h *agentHandlers) UpdateAvatar(c *fiber.Ctx) error {
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

	res, err := h.AgentService.UpdateAvatar(uuid.MustParse(id), payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}

// UpdateAgentByAdmin put handler.
// @Description update agent data by admin with agent id.
// @Summary update agent data
// @Tags Agent
// @Accept json
// @Produce json
// @Param adminId path string true "ID of Admin"
// @Param updateAccount body dto.UpdateAgentByAdmin true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /{adminId}/agent/update [put]
// @Security ApiKeyAuth
func (h *agentHandlers) UpdateAgentByAdmin(c *fiber.Ctx) error {
	payload := new(dto.UpdateAgentByAdmin)
	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	res, err := h.AgentService.UpdateAgentByAdmin(payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}
