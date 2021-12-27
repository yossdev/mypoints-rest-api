package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/configs"
	_ "github.com/yossdev/mypoints-rest-api/docs" // load API Docs files (Swagger)
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/internal/routes"
	_s "github.com/yossdev/mypoints-rest-api/internal/utils/start-server"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	_admin "github.com/yossdev/mypoints-rest-api/src/admins/repositories"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/repositories"
	_product "github.com/yossdev/mypoints-rest-api/src/products/repositories"
	_transaction "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"gorm.io/gorm"
	"log"
)

// @title MyPoints API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService https://swagger.io/terms/

// @contact.name MyPoints Team Support
// @contact.email zenhanprogram@gmail.com

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

	psqlDB := db.NewPsqlClient()
	dbMigrate(psqlDB.DB())

	mongoDB := db.NewMongoClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:     app,
			PsqlDB:  psqlDB,
			MongoDB: mongoDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	//_s.StartServer(app)
	_s.StartServerWithGracefulShutdown(app)
}

// dbMigrate func will auto migrate model struct from record.go in repositories
func dbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&_admin.Admin{},
		&_agent.Agent{},
		&_transaction.TransactionType{},
		&_transaction.TransactionStatus{},
		&_product.Product{},
		&_transaction.Transaction{},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}
