package controllers

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type PositionController struct {
	Usecase *usecase.PositionUsecase
}

func NewPositionController(usecase *usecase.PositionUsecase) *PositionController {
	return &PositionController{Usecase: usecase}
}

func (c *PositionController) CreatePosition(ctx *fiber.Ctx) error {
	var req models.PositionReq
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}
	if req.NamePosition == nil || req.Salary == nil {
		log.Println("Harap isi seluruh kolom")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Harap isi seluruh kolom",
		})
	}
	position, err := c.Usecase.CreatePosition(&req)
	if err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan data posisi",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": fiber.StatusOK,
		"data":    position,
	})
}

func (c *PositionController) DeletePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")
	if positionID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}

	err := c.Usecase.DeletePosition(positionID)
	if err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data posisi",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Data posisi berhasil dihapus",
	})
}

func (c *PositionController) UpdatePosition(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")
	if positionID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}

	var req models.PositionReq
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}

	position, err := c.Usecase.UpdatePosition(positionID, &req)
	if err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memperbarui data posisi",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": fiber.StatusOK,
		"data":    position,
	})
}

func (c *PositionController) GetPositionById(ctx *fiber.Ctx) error {
	positionID := ctx.Params("id")
	if positionID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}

	position, err := c.Usecase.GetPositionById(positionID)
	if err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan data posisi",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": fiber.StatusOK,
		"data":    position,
	})
}
