package repository

import (
	"finalproject_basisdata/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}
//koreksi ini
//seharusnya company tidak berubah
//error
//
// 2023/12/13 04:27:00 D:/belajargolang/finalproject_basisdata/repository/company.go:19 record not found
// [1.426ms] [rows:0] SELECT * FROM `company` WHERE `company`.`company_id` = '3' ORDER BY `company`.`company_id` LIMIT 1
// 2023/12/13 04:27:00 Err:  record not found

func (c *CompanyRepository) GetCompanyInfo(company_id string) (*models.Company, error) {
	var company models.Company
	if err := c.DB.First(&company, company_id).Error; err != nil {
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

func (c *CompanyRepository) UpdateCompanyTopupBalance(company models.Company, balance int) error {

	newBalance := *company.Balance + balance
	if err := c.DB.Model(&company).Update("balance", newBalance).Error; err != nil {
		return err
	}

	transaction := models.Transaction{
		Amount: &balance,
		Note:   "topup saldo perusahaan",
		Type:   "debit",
	}

	if err := c.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (c *CompanyRepository) UpdateCompanyBalanceAfterWithdraw(company models.Company, salary int) error {
	newBalance := *company.Balance - salary
	if err := c.DB.Model(&company).Update("balance", newBalance).Error; err != nil {
		return err
	}

	transaction := models.Transaction{
		Amount: &salary,
		Note:   "pencairan gaji bulanan karyawan",
		Type:   "kredit",
	}

	if err := c.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

// if *company.Balance == 0 {
// 	return nil, errors.New("Balance: 0")
// }
// if err := c.DB.Model(company).Updates(company).Find(company).Error; err != nil {
// 	return nil, err
// }

// updatedBalance := amount + *company.Balance

//TODO: koreksi flow ini gabisa bisa
//y := *company.Balance
