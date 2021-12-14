package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/configs"
	"log"

	_ "github.com/yossdev/mypoints-rest-api/docs" // load API Docs files (Swagger)
)

// @title MyPoints API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService https://swagger.io/terms/

// @contact.name API Support
// @contact.email mypoints@swagger.io

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	SwaggerRoute(app)

	log.Fatal(app.Listen(":8080"))
}

func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	api := a.Group("api")
	v1 := api.Group("/v1")

	// Swagger Docs
	v1.Get("/swagger/*", swagger.Handler)
}
