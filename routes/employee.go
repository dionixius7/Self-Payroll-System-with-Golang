package routes

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(api fiber.Router) {
	// Assuming you have a CompanyRepository instance
	companyRepo := repository.NewCompanyRepository(repository.DB)

	employeeRepo := repository.NewEmployeeRepository(repository.DB)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeRepo, companyRepo)
	employeeController := controllers.NewEmployeeController(employeeUsecase)
	
	// http://localhost:8000/api/admin/company/employee/
	employee := api.Group("/admin/company/employee")
	employee.Post("/", employeeController.CreateEmployeeData)
	employee.Get("/", employeeController.GetAllEmployee)
	employee.Get("/:id", employeeController.GetEmployeeById)
	employee.Patch("/:id", employeeController.UpdateEmployeeData)
	employee.Delete("/:id", employeeController.DeleteEmployee)
	employee.Post("/withdraw/:id", employeeController.WithdrawSalaryEmployee)
}
