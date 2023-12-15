package repository

import (
	"errors"
	"finalproject_basisdata/models"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (c *AdminRepository) CreateAdmin(admin *models.Admin) error {
	return c.DB.Create(admin).Error
}

func (c *AdminRepository) CheckEmailExists(email string) error {
	var admin models.Admin
	var count int64
	if err := c.DB.Model(&admin).Where("email=?", email).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	if count != 0 {
		return fmt.Errorf("Email anda sudah terdaftar: %s", email)
	}
	return nil
}

func (c *AdminRepository) CheckNameExists(name string) error {
	var admin models.Admin
	var count int64
	if err := c.DB.Model(&admin).Where("name=?", name).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	if count != 0 {
		return fmt.Errorf("Nama anda belum dimasukkan.")
	}
	return nil
}

func (c *AdminRepository) CheckPassword(email, password string) error {
	var admin models.Admin
	var count int64
	if err := c.DB.Model(&admin).Where("email=?", email).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	if count != 0 {
		return fmt.Errorf("Email anda sudah terdaftar: %s", email)
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	return nil
}

// fungsi validasi
func validatePassword(password string) error {
	// Check if the password has at least 8 characters
	if len(password) < 8 {
		return fmt.Errorf("Password minimal memiliki 8 karakter")
	}

	hasUppercase := false
	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUppercase = true
			break
		}
	}
	if !hasUppercase {
		return errors.New("Password minimal memiliki 1 huruf kapital")
	}

	hasDigit := false
	for _, char := range password {
		if '0' <= char && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return errors.New("Password minimal memiliki 1 angka")
	}

	hasSpecial := false
	specialChars := "!@#$%^&*()-=_+[]{}|;:'\",.<>/?`~"
	for _, char := range password {
		if strings.ContainsRune(specialChars, char) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return errors.New("Password minimal memiliki 1 karakter spesial")
	}

	return nil
}

//	func (c *AdminRepository) LoginAdmin() {
//		// Compare hashed password
//		err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
//		if err != nil {
//			return fiber.NewError(fiber.StatusBadRequest, "Invalid password")
//		}
//	}
func (c *AdminRepository) GetHashedPassword(email string) (string, error) {
	var admin models.Admin
	if err := c.DB.Model(&admin).Select("password").Where("email = ?", email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("User not found")
		}
		return "", err
	}

	return admin.Password, nil
}

func (c *AdminRepository) CheckEmailLogin(email string) error {
	var admin models.Admin
	var count int64
	if err := c.DB.Model(&admin).Where("email=?", email).Count(&count).Error; err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}
	// if count != 0 {
	// 	return fmt.Errorf("Email anda sudah terdaftar: %s", email)
	// }
	return nil
}
