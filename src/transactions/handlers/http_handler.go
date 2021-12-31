package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

// TransactionHandlers contains method used for the handler
type TransactionHandlers interface {
	CreateTransaction(c *fiber.Ctx) error
}

type transactionHandler struct {
	transactionService entities.Service
}

func NewHttpHandler(s entities.Service) *transactionHandler {
	return &transactionHandler{
		transactionService: s,
	}
}

func (h *transactionHandler) CreateTransaction(c *fiber.Ctx) error {
	return nil
}
