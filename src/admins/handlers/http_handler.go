package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/admins/dto"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
)

// AdminHandlers contains method used for the handler
type AdminHandlers interface {
	SignIn(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	UpdateAvatar(c *fiber.Ctx) error
}

type adminHandlers struct {
	AdminService entities.Service
}

func NewHttpHandler(s entities.Service) AdminHandlers {
	return &adminHandlers{
		AdminService: s,
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
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /admin/login [post]
func (h *adminHandlers) SignIn(c *fiber.Ctx) error {
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

	res, err := h.AdminService.SignIn(payload.ToDomain())
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
// @Failure 400 {object} web.ErrorResp
// @Failure 409 {object} web.ErrorResp
// @Router /admin/signup [post]
func (h *adminHandlers) SignUp(c *fiber.Ctx) error {
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

	res, err := h.AdminService.SignUp(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusConflict, web.DuplicateData, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.AccountCreated, dto.FromDomainAC(res))
}

// UpdateAdmin put handler.
// @Description update admin data by id.
// @Summary update admin data
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin to update"
// @Param updateAccount body dto.UpdateAccount true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /admin/profile/{id} [put]
// @Security ApiKeyAuth
func (h *adminHandlers) UpdateAdmin(c *fiber.Ctx) error {
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

	res, err := h.AdminService.UpdateAdmin(uuid.MustParse(id), payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}

// UpdateAvatar put handler.
// @Description update admin photo profile by id.
// @Summary update admin photo profile
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin to update"
// @Param updateAvatar body dto.UpdateAvatar true "body request"
// @Success 200 {object} dto.AccountUpdated
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /admin/profile/avatar/{id} [put]
// @Security ApiKeyAuth
func (h *adminHandlers) UpdateAvatar(c *fiber.Ctx) error {
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

	res, err := h.AdminService.UpdateAvatar(uuid.MustParse(id), payload.ToDomain())
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainAU(res))
}
