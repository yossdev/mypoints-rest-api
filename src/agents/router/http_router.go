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

	// Private
	v1.Post("/:id/agent", middleware.JwtVerifyToken, agentHandler.SignUp) // sign-up agent by admin
	v1.Get("/profile/:id", middleware.JwtVerifyToken, agentHandler.GetAgent)
	v1.Put("/profile/:id", middleware.JwtVerifyToken, agentHandler.UpdateAgent)
}
