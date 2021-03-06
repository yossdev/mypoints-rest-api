package web

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Meta struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Errors  []string `json:"errors,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func JsonResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	resp := BaseResponse{}
	resp.Meta.Status = statusCode
	resp.Meta.Message = message
	resp.Data = data

	return c.Status(statusCode).JSON(resp)
}

func JsonErrorResponse(c *fiber.Ctx, statusCode int, message error, err error) error {
	resp := BaseResponse{}
	resp.Meta.Status = statusCode
	resp.Meta.Message = message.Error()
	resp.Meta.Errors = []string{err.Error()}

	return c.Status(statusCode).JSON(resp)
}
