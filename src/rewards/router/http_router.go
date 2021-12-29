package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/rewards/handlers"
	"github.com/yossdev/mypoints-rest-api/src/rewards/repositories"
	"github.com/yossdev/mypoints-rest-api/src/rewards/services"
)

type HttpRouter struct {
	web.RouterStruct
}

func NewHttpRoute(r HttpRouter) HttpRouter {
	return r
}

func (r *HttpRouter) GetRoute() {
	rewardPsqlRepository := repositories.NewRewardPsqlRepository(r.PsqlDB)
	rewardService := services.NewRewardService(rewardPsqlRepository)
	rewardHandler := handlers.NewHttpHandler(rewardService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	//	Private - for admin only
	v1.Post("/reward/:id", middleware.JwtVerifyTokenAdmin, rewardHandler.CreateReward)
	v1.Put("/reward/:id/:rewardId", middleware.JwtVerifyTokenAdmin, rewardHandler.UpdateReward)
	v1.Delete("/reward/:id/", middleware.JwtVerifyTokenAdmin, rewardHandler.DeleteReward)
}
