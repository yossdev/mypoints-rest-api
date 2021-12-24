package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/admins/handlers"
	"github.com/yossdev/mypoints-rest-api/src/admins/repositories"
	"github.com/yossdev/mypoints-rest-api/src/admins/services"
)

type HttpRouter struct {
	web.RouterStruct
}

func NewHttpRoute(r HttpRouter) HttpRouter {
	return r
}

func (r *HttpRouter) GetRoute() {
	adminPsqlRepository := repositories.NewAdminPsqlRepository(r.PsqlDB)
	adminService := services.NewAdminService(adminPsqlRepository)
	adminHandler := handlers.NewHttpHandler(adminService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	// Public
	v1.Post("/admin/login", adminHandler.SignIn)
	v1.Post("/admin/signup", adminHandler.SignUp)

	// Private

}
