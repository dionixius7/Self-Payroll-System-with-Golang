package controllers

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/usecase"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	Usecase *usecase.AdminUsecase
}

func NewAdminController(usecase *usecase.AdminUsecase) *AdminController {
	return &AdminController{Usecase: usecase}
}

func (c *AdminController) SingUpAdmin(ctx *fiber.Ctx) error {
	var req models.ReqSignUp
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req",
		})
	}
	err := c.Usecase.SignUpAdmin(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fiber.StatusOK,
		"data":    "Admin berhasil didaftarkan",
	})
}

func (c *AdminController) LoginAdmin(ctx *fiber.Ctx) error {
	var req models.LoginPayload
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req:",
		})
	}
	err := c.Usecase.LoginAdmin(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON("Admin berhasil login")
	// return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": fiber.StatusOK,
	// 	"data":    "Admin berhasil login",
	// })
}
