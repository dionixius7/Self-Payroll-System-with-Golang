package controllers

import (
	//"finalproject_basisdata/models"
	//"finalproject_basisdata/repository"
	"finalproject_basisdata/models"
	"finalproject_basisdata/usecase"
	"log"

	//"strconv"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	Usecase *usecase.CompanyUsecase
}

func NewCompanyController(usecase *usecase.CompanyUsecase) *CompanyController {
	return &CompanyController{Usecase: usecase}
}

func (c *CompanyController) GetCompanyInfo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	company, err := c.Usecase.GetCompanyInfo(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error: Data tidak ditemukan atau Anda salah memasukkan ID",
		})
	}
	if err != nil {
		log.Println("Terjadi kesalahan dalam sistem:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi Kesalahan Dalam Sistem",
		})
	}
	return ctx.JSON(fiber.Map{
		"id":      id,
		"company": company,
	})
}

func (c *CompanyController) CreateCompany(ctx *fiber.Ctx) error {
	var req models.CompanyRequest

	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if req.Name == "" {
		log.Println("Nama perusahaan belum diisi")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama Perusahaan Belum Diisi",
		})
	}
	company, err := c.Usecase.CreateCompany(&req)
	if err != nil {
		log.Println("Failed to create company: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat data perusahaan",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": fiber.StatusOK,
		"data":    company,
	})
}

func (c *CompanyController) UpdateCompany(ctx *fiber.Ctx) error {
	var req models.CompanyRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	id := ctx.Params("id")
	company, err := c.Usecase.UpdateCompany(id, &req)
	if err != nil {
		log.Println("ID perusahaan tidak dapat ditemukan: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Perusahaan Tidak Ditemukan",
		})
	}
	//company, err := c.Usecase.UpdateCompany(id, &req)
	// if err != nil {
	// 	log.Println("Failed to update company: ", err)
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Gagal mengubah data perusahaan",
	// 	})
	// }
	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengubah data perusahaan",
		"data":    company,
	})
}

func (c *CompanyController) TopupBalanceCompany(ctx *fiber.Ctx) error {
	var req models.TopupCompanyBalance
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	id := ctx.Params("id")
	balance, err := c.Usecase.TopupBalanceCompany(id, &req)
	if err != nil {
		log.Println("ID perusahaan tidak dapat ditemukan: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Perusahaan Tidak Ditemukan",
		})
	}
	//company, err := c.Usecase.TopupBalanceCompany(id, &req)
	if err != nil {
		log.Println("Gagal menambahkan saldo perusahaan: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan saldo perusahaan",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Berhasil menambahkan saldo perusahaan",
		"data":    balance,
	})
}
