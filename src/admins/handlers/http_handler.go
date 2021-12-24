package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
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

	// Create a new validator.
	validate := helpers.NewValidator()
	// Validate fields from payload.
	if err := validate.Struct(payload); err != nil {
		// Return, if some fields are not valid.
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	res, err := h.adminService.SignIn(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, web.BadCredential, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Welcome, res)
}

// SignUp post handler.
// @Description create admin account.
// @Summary admins can create from register page
// @Tags Admin
// @Scheme https
// @Accept json
// @Produce json
// @Param signUp body dto.SignUpReq true "body request"
// @Success 201 {object} dto.AccountCreated
// @Router /admin/signup [post]
func (h *adminHandler) SignUp(c *fiber.Ctx) error {
	payload := new(dto.SignUpReq)

	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	// TODO add struct validator

	res, err := h.adminService.SignUp(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusConflict, web.DuplicateData, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.Created, dto.AccountCreated{RowsAffected: res})
}
