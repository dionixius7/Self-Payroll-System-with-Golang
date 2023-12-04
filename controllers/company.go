package controllers

import (
	//"finalproject_basisdata/models"
	//"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	Usecase *usecase.CompanyUsecase
}

func NewCompanyController(usecase *usecase.CompanyUsecase) *CompanyController {
	return &CompanyController{Usecase: usecase}
}

func (c *CompanyController) GetCompanyInfo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		log.Println("Invalid ID:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error: Data tidak ditemukan atau Anda salah memasukkan ID",
		})
	}
	name, err := c.Usecase.GetCompanyInfo(id)
	if err != nil {
		log.Println("Terjadi kesalahan dalam sistem:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi Kesalahan Dalam Sistem",
		})
	}
	return ctx.JSON(fiber.Map{
		"id":   id,
		"name": name,
	})
}

// func GetDetailCompany(c *fiber.Ctx) error {
// 	var company []models.Company

//		repository.DB.Find(&company)
//		return c.JSON(company)
//	}
// func CreateCompany(c *fiber.Ctx) error {
// 	var req models.CompanyRequest
// 	var company models.Company

// 	if err := c.BodyParser(&company); err != nil {
// 		return err
// 	}
// 	if err := repository.DB.Create(&company).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(req)
// }

// func UpdateCompany(c *fiber.Ctx) error {
// 	var req models.CompanyRequest
// 	var company models.Company
// 	id := c.Params("id")
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}
// 	err := repository.DB.Model(models.Company{}).Where("id", id).Debug().Updates(&company).Error
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Tidak dapat memperbarui data: " + err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.StatusOK)
// }

// func TopupBalanceCompany(c *fiber.Ctx) error {
// 	var req models.TopupCompanyBalance
// 	var company models.Company
// 	id := c.Params("id")
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	err := repository.DB.Model(models.Company{}).Where("id", id).Debug().Update(&company).Error
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Tidak dapat memperbarui data: " + err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.StatusOK)
// }
