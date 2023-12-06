package repository

import (
	//"errors"
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

func (c *CompanyRepository) GetCompanyInfo(id string) (*models.Company, error) {
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
	err := c.DB.Model(company).Debug().Updates(company).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CompanyRepository) UpdateCompanyBalance(company *models.Company, balance int) (*models.Company, error) {
	if company.Balance == nil {
		return nil, errors.New("Balance: 0")
	}
	if err := c.DB.Model(company).Updates(company).Find(company).Error; err != nil {
		return nil, err
	}

	// updatedBalance := amount + *company.Balance

	//TODO: koreksi flow ini gabisa bisa cok
	y := *company.Balance
	if err := c.DB.Model(company).Update("Balance", balance+y).Error; err != nil {
		return nil, err
	}
	return company, nil
}
