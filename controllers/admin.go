package controllers

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"finalproject_basisdata/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SignUpAdmin(c *fiber.Ctx) error {
	var request models.ReqSignUp

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//email name password
	if request.Email == "" || request.Name == "" || request.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Isi seluruh form",
		})
	}
	errEmail := usecase.CekEmail(request.Email)
	if errEmail != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errEmail.Error(),
		})
	}
	errName := usecase.CekName(request.Name)
	if errName != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errName.Error(),
		})
	}
	errPass := usecase.CekPassword(request.Password)
	if errPass != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errPass.Error(),
		})
	}
	hashedpassword, err := usecase.HashPassword(request.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newAdmin := models.Admin{
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedpassword,
	}
	if err := repository.DB.Create(&newAdmin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		}) //jika error dalam db maka akan melempar error
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Admin sukses didaftarkan.",
	})
}

func LoginAdmin(c *fiber.Ctx) error {
	var loginPayload models.LoginPayload
	var admins models.Admin

	if err := c.BodyParser(&loginPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	email := loginPayload.Email
	pass := loginPayload.Password
	if err := repository.DB.First(&admins, "email", email).Error; err != nil {
		log.Print("Data tidak ditemukan")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	var hashedpassword string
	if err := repository.DB.Select("password").First(&admins, "email", email).Scan(&hashedpassword).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}
	match := usecase.CheckPasswordHash(pass, hashedpassword)
	if match {
		//MASIH GATAU INI APA
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login successful",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password anda salah!",
		})
	}
}
