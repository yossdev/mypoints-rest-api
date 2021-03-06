package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/transactions/dto"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

// TransactionHandlers contains method used for the handler
type TransactionHandlers interface {
	Claims(c *fiber.Ctx) error
	UpdateClaims(c *fiber.Ctx) error
	Redeem(c *fiber.Ctx) error
	CallbackXendit(c *fiber.Ctx) error
}

type transactionHandlers struct {
	TransactionService entities.Service
}

func NewHttpHandler(s entities.Service) TransactionHandlers {
	return &transactionHandlers{
		TransactionService: s,
	}
}

// Claims post handler.
// @Description create claims transaction by agents.
// @Summary agent can create claims transaction
// @Tags Transaction
// @Scheme https
// @Accept json
// @Produce json
// @Param id path string true "ID of Agent"
// @Param newClaims body dto.ClaimsReq true "body request"
// @Success 201 {object} dto.TransactionRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 424 {object} web.ErrorResp
// @Router /{id}/transactions/claims [post]
// @Security ApiKeyAuth
func (h *transactionHandlers) Claims(c *fiber.Ctx) error {
	payload := new(dto.ClaimsReq)
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

	res, err := h.TransactionService.Claims(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusFailedDependency, web.Failed, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.ClaimsTransactionCreated, dto.FromDomain(res))
}

// UpdateClaims put handler.
// @Description update claims transaction status by admins.
// @Summary admins can update claims transaction from agent
// @Tags Transaction
// @Scheme https
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param updateClaims body dto.UpdateClaimsReq true "body request"
// @Success 200 {object} dto.TransactionRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Router /admin/{id}/transactions/claims [put]
// @Security ApiKeyAuth
func (h *transactionHandlers) UpdateClaims(c *fiber.Ctx) error {
	payload := new(dto.UpdateClaimsReq)
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

	res, err := h.TransactionService.ClaimsStatus(payload.ID, payload.Status)
	if err != nil || res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomain(res))
}

// Redeem post handler.
// @Description create redeem transaction by agents.
// @Summary agent can create redeem transaction
// @Tags Transaction
// @Scheme https
// @Accept json
// @Produce json
// @Param id path string true "ID of Agent"
// @Param newRedeem body dto.RedeemReq true "body request"
// @Success 201 {object} dto.TransactionRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 424 {object} web.ErrorResp
// @Router /{id}/transactions/redeem [post]
// @Security ApiKeyAuth
func (h *transactionHandlers) Redeem(c *fiber.Ctx) error {
	payload := new(dto.RedeemReq)
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

	res, err := h.TransactionService.Redeem(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusFailedDependency, web.Failed, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.RedeemTransactionCreated, dto.FromDomain(res))
}

// CallbackXendit post handler called by xendit after admin paid the invoice
// @Description callback for invoice xendit
// @Summary callback endpoint
// @Tags Transaction
// @Scheme https
// @Accept json
// @Produce json
// @Param invoiceCallback body dto.InvoiceCallback true "body request"
// @Success 200 {string} string web.RedeemTransactionCompleted
// @Router /redeem/callback [post]
func (h *transactionHandlers) CallbackXendit(c *fiber.Ctx) error {
	token := c.Get("x-callback-token")

	payload := new(dto.InvoiceCallback)
	if err := c.BodyParser(payload); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	if token != viper.GetString("X_Callback_Token") {
		return web.JsonErrorResponse(c, fiber.StatusNotAcceptable, web.InvalidToken, web.InvalidToken)
	}

	if err := h.TransactionService.CallbackXendit(payload.ToDomain()); err != nil {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, err)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.RedeemTransactionCompleted, nil)
}
