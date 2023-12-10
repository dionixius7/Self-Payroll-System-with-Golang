package routes

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupCompanyRoutes(api fiber.Router) {
	companyRepo := repository.NewCompanyRepository(repository.DB)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyController := controllers.NewCompanyController(companyUsecase)

	//http://localhost:8000/api/admin/company/
	company := api.Group("/admin/company")
	company.Get("/:id", companyController.GetCompanyInfo)
	company.Post("/", companyController.CreateCompany)
	company.Patch("/:id", companyController.UpdateCompany)
	company.Patch("/topup/:id", companyController.TopupBalanceCompany)
}
