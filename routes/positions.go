package routes

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupPositionRoutes(api fiber.Router) {
	positionRepo := repository.NewPositionRepository(repository.DB)
	positionUsecase := usecase.NewPositionUsecase(positionRepo)
	positionController := controllers.NewPositionController(positionUsecase)

	//http://localhost:8000/api/admin/company/position/
	position := api.Group("/positions")
	position.Get("/", positionController.GetAllPositionData)
	position.Get("/:id", positionController.GetPositionById)
	position.Post("/", positionController.CreatePosition)
	position.Delete("/:id", positionController.DeletePosition)
	position.Patch("/:id", positionController.UpdatePosition)
}
