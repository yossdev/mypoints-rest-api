package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
)

// ProductHandlers contains method used for the handler
type ProductHandlers interface {
	CreateProduct(c *fiber.Ctx) error
}

type productHandler struct {
	productService entities.Service
}

func NewHttpHandler(s entities.Service) *productHandler {
	return &productHandler{
		productService: s,
	}
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	return nil
}
