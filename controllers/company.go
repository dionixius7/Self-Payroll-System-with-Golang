package controllers

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	Usecase *usecase.CompanyUsecase
}

func NewCompanyController(usecase *usecase.CompanyUsecase) *CompanyController {
	return &CompanyController{Usecase: usecase}
}
func (c *CompanyController) GetAllCompany(ctx *fiber.Ctx) error {
	companies, err := c.Usecase.GetAllCompany()
	if err != nil {
		log.Println("Error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan seluruh data perusahaan yang tersimpan",
		})
	}
	return ctx.JSON(companies)
}

func (c *CompanyController) GetCompanyInfo(ctx *fiber.Ctx) error {
	company_id := ctx.Params("company_id")
	company, err := c.Usecase.GetCompanyInfo(company_id)
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
	return ctx.JSON(company)
	// return ctx.JSON(fiber.Map{
	// 	"company_id": company_id,
	// 	"company":    company,
	// })
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
	return ctx.JSON(company)

	// return ctx.JSON(fiber.Map{
	// 	"message": fiber.StatusOK,
	// 	"data":    company,
	// })
}

func (c *CompanyController) UpdateCompany(ctx *fiber.Ctx) error {
	var req models.CompanyRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	company_id := ctx.Params("company_id")
	company, err := c.Usecase.UpdateCompany(company_id, &req)
	if err != nil {
		log.Println("ID perusahaan tidak dapat ditemukan: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Perusahaan Tidak Ditemukan",
		})
	}
	return ctx.JSON(company)

	// return ctx.JSON(fiber.Map{
	// 	"message": "Berhasil mengubah data perusahaan",
	// 	"data":    company,
	// })
}

func (c *CompanyController) TopupBalanceCompany(ctx *fiber.Ctx) error {
	var req models.TopupCompanyBalance
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if req.Balance == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Saldo harus lebih besar dari nol",
		})
	}
	company_id := ctx.Params("company_id")
	balance, err := c.Usecase.TopupBalanceCompany(company_id, req)
	if err != nil {
		log.Println("ID perusahaan tidak dapat ditemukan: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Perusahaan Tidak Ditemukan",
		})
	}
	return ctx.JSON(balance)
	// return ctx.JSON(fiber.Map{
	// 	"message": "Berhasil menambahkan saldo perusahaan",
	// 	"data":    balance,
	// })
}
