package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
)

type transactionHandler struct {
	transactionService entities.Service
}

func NewHttpHandler(s entities.Service) *transactionHandler {
	return &transactionHandler{
		transactionService: s,
	}
}

func (h *transactionHandler) GetTransactions(c *fiber.Ctx) error {
	return nil
}

func (h *transactionHandler) GetTransactionDetail(c *fiber.Ctx) error {
	return nil
}

func (h *transactionHandler) GetTransactionsAdmin(c *fiber.Ctx) error {
	return nil
}
