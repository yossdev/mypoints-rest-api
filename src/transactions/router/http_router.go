package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	_admin "github.com/yossdev/mypoints-rest-api/src/admins/repositories"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/repositories"
	_reward "github.com/yossdev/mypoints-rest-api/src/rewards/repositories"
	"github.com/yossdev/mypoints-rest-api/src/transactions/handlers"
	"github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"github.com/yossdev/mypoints-rest-api/src/transactions/services"
)

type HttpRouter struct {
	web.RouterStruct
}

func NewHttpRoute(r HttpRouter) HttpRouter {
	return r
}

func (r *HttpRouter) GetRoute() {
	transactionPsqlRepository := repositories.NewTransactionPsqlRepository(r.PsqlDB)
	agentPsqlRepository := _agent.NewAgentPsqlRepository(r.PsqlDB)
	adminPsqlRepository := _admin.NewAdminPsqlRepository(r.PsqlDB)
	rewardPsqlRepository := _reward.NewRewardPsqlRepository(r.PsqlDB)
	transactionService := services.NewTransactionService(transactionPsqlRepository, agentPsqlRepository, adminPsqlRepository, rewardPsqlRepository)
	transactionHandler := handlers.NewHttpHandler(transactionService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	admin := v1.Group("/admin/:id/transactions") // param for agent id
	agent := v1.Group("/:id/transactions")       // param for agent id

	// Public
	v1.Post("/redeem/callback", transactionHandler.CallbackXendit)

	// Private
	// Agent API only for agent
	agent.Post("/claims", middleware.JwtVerifyTokenAgent, transactionHandler.Claims)
	agent.Post("/redeem", middleware.JwtVerifyTokenAgent, transactionHandler.Redeem)

	// Admin API only for admin
	admin.Put("/claims/:transactionId", middleware.JwtVerifyTokenAdmin, transactionHandler.UpdateClaims)
}
