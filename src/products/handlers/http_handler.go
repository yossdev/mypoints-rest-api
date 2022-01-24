package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/products/dto"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
)

// ProductHandlers contains method used for the handler
type ProductHandlers interface {
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type productHandlers struct {
	ProductService entities.Service
}

func NewHttpHandler(s entities.Service) ProductHandlers {
	return &productHandlers{
		ProductService: s,
	}
}

// CreateProduct post handler.
// @Description create product by admins.
// @Summary admins can create product
// @Tags Product
// @Scheme https
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param newProduct body dto.NewProduct true "body request"
// @Success 201 {object} dto.ProductRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 424 {object} web.ErrorResp
// @Router /product/{id} [post]
// @Security ApiKeyAuth
func (h *productHandlers) CreateProduct(c *fiber.Ctx) error {
	payload := new(dto.NewProduct)
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

	res, err := h.ProductService.CreateProduct(payload.ToDomain())
	if err != nil {
		return web.JsonErrorResponse(c, fiber.StatusFailedDependency, web.Failed, err)
	}

	return web.JsonResponse(c, fiber.StatusCreated, web.ProductCreated, dto.FromDomainRA(res))
}

// UpdateProduct put handler.
// @Description update product data by id.
// @Summary update product data
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param productId path int true "ID of Product"
// @Param updateProduct body dto.UpdateProduct true "body request"
// @Success 200 {object} dto.ProductRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 422 {object} web.ErrorResp
// @Router /product/{id}/{productId} [put]
// @Security ApiKeyAuth
func (h *productHandlers) UpdateProduct(c *fiber.Ctx) error {
	params := c.Params("productId")
	productId, convErr := helpers.StringToUint32(params)
	if convErr != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnprocessableEntity, web.CannotProcess, convErr)
	}

	payload := new(dto.UpdateProduct)
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

	res := h.ProductService.UpdateProduct(productId, payload.ToDomain())
	if res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, fiber.ErrBadRequest)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainRA(res))
}

// DeleteProduct delete handler.
// @Description soft delete product data by id.
// @Summary soft delete product data
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "ID of Admin"
// @Param productId path int true "ID of Product"
// @Success 200 {object} dto.ProductRes
// @Failure 400 {object} web.ErrorResp
// @Failure 401 {object} web.ErrorResp
// @Failure 422 {object} web.ErrorResp
// @Router /product/{id}/{productId} [delete]
// @Security ApiKeyAuth
func (h *productHandlers) DeleteProduct(c *fiber.Ctx) error {
	params := c.Params("productId")
	productId, convErr := helpers.StringToUint32(params)
	if convErr != nil {
		return web.JsonErrorResponse(c, fiber.StatusUnprocessableEntity, web.CannotProcess, convErr)
	}

	res := h.ProductService.DeleteProduct(productId)
	if res == 0 {
		return web.JsonErrorResponse(c, fiber.StatusBadRequest, web.BadRequest, fiber.ErrBadRequest)
	}

	return web.JsonResponse(c, fiber.StatusOK, web.Success, dto.FromDomainRA(res))
}
