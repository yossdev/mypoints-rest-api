package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

// TransactionHandlers contains method used for the handler
type TransactionHandlers interface {
	CreateTransaction(c *fiber.Ctx) error
}

type transactionHandlers struct {
	TransactionService entities.Service
}

func NewHttpHandler(s entities.Service) TransactionHandlers {
	return &transactionHandlers{
		TransactionService: s,
	}
}

func (h *transactionHandlers) CreateTransaction(c *fiber.Ctx) error {
	return nil
}
