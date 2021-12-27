package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
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
	transactionService := services.NewTransactionService(transactionPsqlRepository)
	transactionHandler := handlers.NewHttpHandler(transactionService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	admin := v1.Group("/:admin/transactions") // admin param for agent id
	agent := v1.Group("/:agent/transactions") // agent param for agent id

	// Private
	// Agent API only for agent
	agent.Get("/", middleware.JwtVerifyToken, transactionHandler.GetTransactions)
	agent.Get("/:id", middleware.JwtVerifyToken, transactionHandler.GetTransactionDetail)

	// Admin API only for admin
	admin.Get("/", middleware.JwtVerifyToken, transactionHandler.GetTransactionsAdmin)
}
