package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
)

type RouterStruct struct {
	Web     *fiber.App
	PsqlDB  db.PsqlDB
	MongoDB db.MongoDB
}
