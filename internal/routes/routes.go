package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"log"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (r *RouterStruct) GetRoutes() {
	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	v1.Use(logger.New(), cors.New())

	// base path
	v1.Get("/", func(c *fiber.Ctx) error {
		return web.JsonResponse(c, fiber.StatusOK, "MyPoints API's V1", nil)
	})

	// Swagger Docs
	v1.Get("/swagger/*", swagger.Handler)

	//webRouterConfig := web.RouterStruct{
	//	Web: r.Web,
	//}

	// registering route from another modules
	// Agent Route

	// handling 404 error
	v1.Use(func(c *fiber.Ctx) error {
		return web.JsonResponse(c, fiber.StatusNotFound, "Sorry can't find that!", nil)
	})
}
