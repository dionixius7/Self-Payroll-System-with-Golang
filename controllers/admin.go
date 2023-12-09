package controllers

// import (
// 	"finalproject_basisdata/models"
// 	"finalproject_basisdata/usecase"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// type AdminController struct {
// 	Usecase *usecase.AdminUsecase
// }

// func NewAdminController(usecase *usecase.AdminUsecase) *AdminController {
// 	return &AdminController{Usecase: usecase}
// }

// func (c *AdminController) SignUpAdmin(ctx *fiber.Ctx) error {
// 	var req models.ReqSignUp
// 	if err := ctx.BodyParser(&req); err != nil {
// 		log.Println("Invalid req body: ", err)
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body",
// 		})
// 	}
// 	if req.Email == "" || req.Name == "" || req.Password == "" {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Tolong isi seluruh form",
// 		})
// 	}
// 	//TODO: koreksi flow checking ini dari ctrl - usecase - repo
// 	errEmail := c.Usecase.CekEmail(&req)
// 	if errEmail != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Kolom email sudah terdaftar",
// 		})
// 	}
// 	errPassword := c.Usecase.CekPassword(req.Password)
// 	if errPassword != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Kolom password harus memiliki minimal 8 karakter, 1 huruf besar, 1 angka dan 1 spesial karakter",
// 		})
// 	}
// 	errName := c.Usecase.CekName(req.Name)
// 	if errName != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Kolom nama belum terisi",
// 		})
// 	}
// 	hashedPassword, err := c.Usecase.HashPassword(req.Password)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	admin := models.Admin{
// 		Name:     req.Name,
// 		Email:    req.Email,
// 		Password: hashedPassword,
// 	}
// 	admin, err := c.Usecase.SignUpAdmin(req)
// 	if err != nil {
// 		log.Println("Gagal membuat data admin: ", err)
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Gagal menambahkan admin baru",
// 		})
// 	}
// 	// You can use hashedPassword here if needed
// 	return ctx.JSON(fiber.Map{
// 		"status": fiber.StatusOK,
// 		"data":   admin,
// 	})
// }

// // func LoginAdmin(c *fiber.Ctx) error {
// // 	var loginPayload models.LoginPayload
// // 	var admins models.Admin

// // 	if err := c.BodyParser(&loginPayload); err != nil {
// // 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// // 			"message": err.Error(),
// // 		})
// // 	}

// // 	email := loginPayload.Email
// // 	pass := loginPayload.Password
// // 	if err := repository.DB.First(&admins, "email", email).Error; err != nil {
// // 		log.Print("Data tidak ditemukan")
// // 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// // 			"message": "Data tidak ditemukan",
// // 		})
// // 	}

// // 	var hashedpassword string
// // 	if err := repository.DB.Select("password").First(&admins, "email", email).Scan(&hashedpassword).Error; err != nil {
// // 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// // 			"message": "Data tidak ditemukan",
// // 		})
// // 	}
// // 	match := usecase.CheckPasswordHash(pass, hashedpassword)
// // 	if match {
// // 		//MASIH GATAU INI APA
// // 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// // 			"message": "Login successful",
// // 		})
// // 	} else {
// // 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// // 			"message": "Password anda salah!",
// // 		})
// // 	}
// // }
