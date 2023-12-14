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
	company := api.Group("/company")
	company.Get("/", companyController.GetAllCompany)
	company.Get("/:company_id", companyController.GetCompanyInfo)
	company.Post("/", companyController.CreateCompany)
	company.Patch("/:company_id", companyController.UpdateCompany)
	company.Patch("/topup/:company_id", companyController.TopupBalanceCompany)
}
