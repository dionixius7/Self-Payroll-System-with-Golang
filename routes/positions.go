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

	position := api.Group("/company/position")
	position.Post("/", positionController.CreatePosition)
	position.Get("/:id", positionController.GetPositionById)
	position.Delete("/:id", positionController.DeletePosition)
	position.Patch("/:id", positionController.UpdatePosition)
}
