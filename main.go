package main

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repository.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	company := api.Group("/company")

	companyRepo := repository.NewCompanyRepository(repository.DB)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyController := controllers.NewCompanyController(companyUsecase)

	company.Get("/:id", companyController.GetCompanyInfo)

	company.Post("/", companyController.CreateCompany)
	company.Patch("/:id", companyController.UpdateCompany)
	company.Patch("/topup/:id", companyController.TopupBalanceCompany)
	app.Listen(":8000")
}
