package router

import (
	"github.com/yossdev/mypoints-rest-api/internal/middleware"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"github.com/yossdev/mypoints-rest-api/src/products/handlers"
	"github.com/yossdev/mypoints-rest-api/src/products/repositories"
	"github.com/yossdev/mypoints-rest-api/src/products/services"
)

type HttpRouter struct {
	web.RouterStruct
}

func NewHttpRoute(r HttpRouter) HttpRouter {
	return r
}

func (r *HttpRouter) GetRoute() {
	productPsqlRepository := repositories.NewProductPsqlRepository(r.PsqlDB)
	productService := services.NewProductService(productPsqlRepository)
	productHandler := handlers.NewHttpHandler(productService)

	api := r.Web.Group("api")
	v1 := api.Group("/v1")

	//	Private - for admin only
	v1.Post("/product/:id", middleware.JwtVerifyTokenAdmin, productHandler.CreateProduct)
	v1.Put("/product/:id/:productId", middleware.JwtVerifyTokenAdmin, productHandler.UpdateProduct)
	v1.Delete("/product/:id/:productId", middleware.JwtVerifyTokenAdmin, productHandler.DeleteProduct)
}
