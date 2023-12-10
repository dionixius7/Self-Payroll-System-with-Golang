package main

import (
	"finalproject_basisdata/repository"
	"finalproject_basisdata/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repository.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")

	routes.SetupCompanyRoutes(api)
	routes.SetupEmployeeRoutes(api)
	routes.SetupAdminRoutes(api)
	routes.SetupPositionRoutes(api)
	routes.SetupTransactionRoutes(api)
	app.Listen(":8000")
}
