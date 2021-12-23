package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/agents/handlers"
	"github.com/yossdev/mypoints-rest-api/src/agents/repositories"
	"github.com/yossdev/mypoints-rest-api/src/agents/services"
)

type HttpRouter struct {
	web.RouterStruct
}

func NewHttpRoute(r HttpRouter) HttpRouter {
	return r
}

func (r *HttpRouter) GetRoute() {
	agentPsqlRepository := repositories.NewAgentPsqlRepository(r.PsqlDB)
	agentService := services.NewAgentService(agentPsqlRepository)
	agentHandler := handlers.NewHttpHandler(agentService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	// Public
	v1.Post("/login", agentHandler.SignIn)

	// Custom JWT Middleware
	v1.Use(middleware.JwtVerifyToken)

	// Private
	v1.Get("/profile/:id", agentHandler.GetAgent)
	v1.Post("/:id/agent", middleware.JwtVerifyToken, agentHandler.SignUp)
}
