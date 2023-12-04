package usecase

import (
	"errors"
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func CekEmail(email string) error {
	var admin models.Admin
	var count int64
	if err := repository.DB.Model(&admin).Where("email=?", email).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	if count != 0 {
		return fmt.Errorf("Email anda sudah terdaftar.", email)
	}
	return nil
}

func CekName(name string) error {
	// var users models.Users
	var admin models.Admin
	var count int64
	if err := repository.DB.Model(&admin).Where("name=?", name).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	if count != 0 {
		return fmt.Errorf("Nama anda belum dimasukkan.")
	}
	return nil
}

func CekPassword(password string) error {
	const (
		minLength      = 8
		minUppercase   = 1
		minDigits      = 1
		minSpecialChar = 1
	)

	if len(password) < minLength {
		return errors.New("Password kurang panjang!")
	}
	uppercaseCount := len(regexp.MustCompile(`[^A-Z]`).ReplaceAllString(password, ""))
	if uppercaseCount < minUppercase {
		return errors.New("Masukkan minimal 1 huruf kapital pada password!")
	}
	digitCount := len(regexp.MustCompile(`[^0-9]`).ReplaceAllString(password, ""))
	if digitCount < minDigits {
		return errors.New("Masukkan minimal 1 angka pada password!")
	}
	specialCharCount := len(regexp.MustCompile(`[a-zA-Z0-9]`).ReplaceAllString(password, ""))
	if specialCharCount < minSpecialChar {
		return errors.New("Masukkan minimal 1 karakter spesial pada password!")
	}
	return nil
}
