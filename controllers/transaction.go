package controllers

import (
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	Usecase *usecase.TransactionUsecase
}

func NewTransactionController(usecase *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{Usecase: usecase}
}

func (c *TransactionController) GetInfoTransaction(ctx *fiber.Ctx) error {
	transactions, err := c.Usecase.GetInfoTransaction()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan data transaksi",
		})
	}
	return ctx.JSON(transactions)

	// return ctx.JSON(fiber.Map{
	// 	"message": fiber.StatusOK,
	// 	"data":    transactions,
	// })
}
