package routes

import (
	"finalproject_basisdata/controllers"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(api fiber.Router) {
	transactionRepo := repository.NewTransactionRepository(repository.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionUsecase)
	
	//http://localhost:8000/api/admin/company/transaction/
	transaction := api.Group("/admin/company/transaction")
	transaction.Get("/", transactionController.GetInfoTransaction)
}
