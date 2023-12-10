package routes

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(api fiber.Router) {
	adminRepo := repository.NewAdminRepository(repository.DB)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)
	adminController := controllers.NewAdminController(adminUsecase)

	//Signup
	//http://localhost:8000/api/admin/signup/
	admin := api.Group("/admin/signup")
	admin.Post("/", adminController.SingUpAdmin)

	//login
	//http://localhost:8000/api/admin/login/
	admin2 := api.Group("/admin/login")
	admin2.Post("/", adminController.LoginAdmin)

}
