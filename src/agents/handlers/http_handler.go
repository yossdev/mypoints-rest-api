package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/agents/dto"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
)

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
// @Tags Agents
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

	// TODO add struct validator

	res, err := h.agentService.SignIn(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, web.BadCredential, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Welcome, res)
}

// GetAgent get handler.
func (h *agentHandler) GetAgent(c *fiber.Ctx) error {
	id := c.Params("id")
	agent, err := h.agentService.GetAgent(uuid.MustParse(id))
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusForbidden, web.Forbidden, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomain(agent))
}
