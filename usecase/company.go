package usecase

import (
	"errors"
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"log"
	//"github.com/gofiber/fiber/v2"
)

type CompanyUsecase struct {
	Repo *repository.CompanyRepository
}

func NewCompanyUsecase(repo *repository.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{Repo: repo}
}

func (c *CompanyUsecase) GetCompanyInfo(id int) (*models.Company, error) {
	company, err := c.Repo.GetCompanyInfo(id)
	if err != nil {
		log.Println("Terdapat error dalam sistem:", err)
		return nil, err
	}
	return company, nil
}

func (c *CompanyUsecase) CreateCompany(req *models.CompanyRequest) (*models.Company, error) {
	if req.Name == "" {
		return nil, errors.New("Silakan isi seluruh field")
	}
	company := &models.Company{
		Name:    req.Name,
		Balance: &req.Balance,
	}
	if err := c.Repo.CreateCompany(company); err != nil {
		log.Println("Tidak dapat membuat data perusahaan:", err)
		return nil, err
	}
	return company, nil
}

func (c *CompanyUsecase) UpdateCompany(id int, req *models.CompanyRequest) (*models.Company, error) {
	if req.Name == "" || req.Balance == 0 {
		return nil, errors.New("Silakan isi seluruh field")
	}
	company := &models.Company{
		ID:      id,
		Name:    req.Name,
		Balance: &req.Balance,
	}
	if err := c.Repo.UpdateCompany(company); err != nil {
		log.Println("Terjadi kesalahan saat mengubah data perusahaan: ", err)
		return nil, err
	}
	return company, nil
}

func (c *CompanyUsecase) TopupBalanceCompany(id int, req *models.TopupCompanyBalance) (*models.Company, error) {
	company, err := c.Repo.GetCompanyInfo(id)
	if err != nil {
		log.Println("Tidak dapat menemukan data perusahaan: ", err)
		return nil, err
	}

	updatedCompany, err := c.Repo.UpdateCompanyBalance(company, req.Balance)
	if err != nil {
		log.Println("Tidak dapat menemukan data perusahaan: ", err)
		return nil, err
	}

	return updatedCompany, nil
}

// company, err := c.Repo.GetCompanyInfo(id)
// if err != nil {
// 	log.Println("Tidak dapat menemukan data perusahaan: ", err)
// 	return nil, err
// }
// y
// company.Balance := company.Balance + req.Balance
// if err := c.Repo.UpdateCompanyBalance(company); err != nil {
// 	log.Println("Terjadi kesalahan saat menambahkan saldo ke perusahaan: ", err)
// 	return nil, err
// }
// return company, nil
