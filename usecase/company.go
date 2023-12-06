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

func (c *CompanyUsecase) GetCompanyInfo(id string) (*models.Company, error) {
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

func (c *CompanyUsecase) UpdateCompany(id string, req *models.CompanyRequest) (*models.Company, error) {
	// if req.Name == "" || req.Balance == 0 {
	// 	return nil, errors.New("Silakan isi seluruh field")
	// }
	// company := &models.Company{
	// 	// ID:      id,
	// 	Name:    req.Name,
	// 	Balance: &req.Balance,
	// }
	company, err := c.Repo.GetCompanyInfo(id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		company.Name = req.Name
	}
	if req.Balance != 0 {
		company.Balance = &req.Balance
	}
	if err := c.Repo.UpdateCompany(company); err != nil {
		// log.Println("Terjadi kesalahan saat mengubah data perusahaan: ", err)
		return nil, err
	}
	return company, nil
}

func (c *CompanyUsecase) TopupBalanceCompany(id string, req *models.TopupCompanyBalance) (*models.Company, error) {
	company, err := c.Repo.GetCompanyInfo(id)
	if err != nil {
		// log.Println("Tidak dapat menemukan data perusahaan: ", err)
		return nil, err
	}

	updatedCompany, err := c.Repo.UpdateCompanyBalance(company, req.Balance)
	if err != nil {
		return nil, err
	}

	return updatedCompany, nil
}
