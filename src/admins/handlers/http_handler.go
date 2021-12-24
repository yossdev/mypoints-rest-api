package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/admins/dto"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
)

type adminHandler struct {
	adminService entities.Service
}

func NewHttpHandler(s entities.Service) *adminHandler {
	return &adminHandler{
		adminService: s,
	}
}

// SignIn post handler.
// @Description check admins by checking given email and password.
// @Summary check admins by given email return jwt token if successfully signIn
// @Tags Admin
// @Scheme https
// @Accept json
// @Produce json
// @Param signIn body dto.SignInReq true "body request"
// @Success 200 {object} auth.Token
// @Router /admin/login [post]
func (h *adminHandler) SignIn(c *fiber.Ctx) error {
	payload := new(dto.SignInReq)

	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	// TODO add struct validator

	res, err := h.adminService.SignIn(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, web.BadCredential, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Welcome, res)
}
