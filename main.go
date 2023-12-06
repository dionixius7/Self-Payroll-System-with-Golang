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

	employeeRepo := repository.NewEmployeeRepository(repository.DB)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeUsecase)
	employee := api.Group("/employee")
	employee.Post("/", employeeController.CreateEmployeeData)
	employee.Get("/", employeeController.GetAllEmployee)
	employee.Get("/:id", employeeController.GetEmployeeById)
	employee.Patch("/:id", employeeController.UpdateEmployeeData)
	employee.Delete("/:id", employeeController.DeleteEmployee)

	app.Listen(":8000")
}
