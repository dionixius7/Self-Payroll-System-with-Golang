package repository

// import (
// 	"finalproject_basisdata/models"

// 	"gorm.io/gorm"
// )

// type AdminRepository struct {
// 	DB *gorm.DB
// }

// func NewAdminRepository(db *gorm.DB) *AdminRepository {
// 	return &AdminRepository{DB: db}
// }

// func (c *AdminRepository) SignUpAdmin(admin *models.Admin) error {
// 	return c.DB.Create(admin).Error
// }

// // koreksi ini seharusnya tidak seperti ini
// func (c *AdminRepository) CekEmail(email string) (*models.Admin, error) {
// 	var admin models.Admin
// 	result := c.DB.Where("email = ?", email).First(&admin)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &admin, nil
// }

// func (c *AdminRepository) CekName(name string) (*models.Admin, error) {
// 	var admin models.Admin
// 	result := c.DB.Where("name = ?", name).First(&admin)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &admin, nil
// }

// // cek ini juga
// func (c *AdminRepository) CekPassword(password string) (*models.Admin, error) {
// 	var admin models.Admin
// 	result := c.DB.Where("password = ?", password).First(&admin)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &admin, nil
// }
