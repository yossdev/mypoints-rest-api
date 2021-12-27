package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
)

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
