package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	_adminRoute "github.com/yossdev/mypoints-rest-api/src/admins/router"
	_agentRoute "github.com/yossdev/mypoints-rest-api/src/agents/router"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	return r
}

func (r *RouterStruct) GetRoutes() {
	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	// Fiber middleware
	v1.Use(logger.New(), cors.New())

	// custom middleware
	v1.Use(middleware.NewLogMongo(r.MongoDB).LogReqRes)

	// base path
	v1.Get("/", func(c *fiber.Ctx) error {
		return web.JsonResponse(c, fiber.StatusOK, "MyPoints API's V1", nil)
	})

	// Swagger Docs
	v1.Get("/swagger/*", swagger.Handler)

	webRouterConfig := web.RouterStruct{
		Web:     r.Web,
		PsqlDB:  r.PsqlDB,
		MongoDB: r.MongoDB,
	}

	// registering route from another modules
	// Admin Route
	adminRouterStruct := _adminRoute.HttpRouter{
		RouterStruct: webRouterConfig,
	}
	adminRouter := _adminRoute.NewHttpRoute(adminRouterStruct)
	adminRouter.GetRoute()

	// Agent Route
	agentRouterStruct := _agentRoute.HttpRouter{
		RouterStruct: webRouterConfig,
	}
	agentRouter := _agentRoute.NewHttpRoute(agentRouterStruct)
	agentRouter.GetRoute()

	// handling 404 error
	v1.Use(func(c *fiber.Ctx) error {
		return web.JsonResponse(c, fiber.StatusNotFound, "Sorry can't find that!", nil)
	})
}
