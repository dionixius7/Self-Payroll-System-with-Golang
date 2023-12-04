package repository

import (
	"errors"
	"finalproject_basisdata/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (c *CompanyRepository) GetCompanyInfo(id int) (*models.Company, error) {
	var company models.Company
	if err := c.DB.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *CompanyRepository) CreateCompany(company *models.Company) error {
	if err := c.DB.Create(company).Error; err != nil {
		return err
	}
	return nil
}

func (c *CompanyRepository) UpdateCompany(company *models.Company) error {
	if err := c.DB.Model(&models.Company{}).Where("id = ?", company.ID).Updates(map[string]interface{}{
		"Name":    company.Name,
		"Balance": company.Balance,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (c *CompanyRepository) UpdateCompanyBalance(company *models.Company, balance int) (*models.Company, error) {
	if company.Balance == nil {
		return nil, errors.New("Saldo perusahaan kosong")
	}
	updatedBalance := *company.Balance + balance

	if err := c.DB.Model(company).Updates(map[string]interface{}{
		"Balance": updatedBalance,
		// "Note":    "Topup saldo perusahaan",
	}).Error; err != nil {
		return nil, err
	}

	updatedCompany, err := c.GetCompanyInfo(company.ID)
	if err != nil {
		return nil, err
	}

	return updatedCompany, nil
}

// func (c *CompanyRepository) UpdateCompanyBalance(company *models.Company) error {
// 	company, err := c.DB.Get(id)
// 	y := company.Balance
// 	if err := c.DB.Model(&models.Company{}).Where("id = ?", company.ID).Error; err != nil {
// 		return nil,err
// 	}
// 	if err := c.DB.Model(&models.Company{}).Where("id = ?", company.ID).Updates("Balance", y + balance).Error; err != nil {
// 		return nil, err
// 	}
// 	if err := c.DB.Model(&models.Company{}).Where("id = ?", company.ID).Updates(map[string]interface{}{
// 		Amount: balance,
// 		Note: "Topup saldo perusahaan",
// 	}).Error; err != nil {
// 		return nil, err
// 	}
// 	return company, err

// }
// if err := c.DB.Save(company).Error; err != nil {
// 	return err
// }
// return nil
