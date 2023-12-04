package usecase

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"log"
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

// type companyUsecase struct {
// 	companyRepo models.CompanyRepository
// }

// func NewCompanyUsecase(repo models.CompanyRepository) models.CompanyUsecase {
// 	return &companyUsecase{companyRepo: repo}
// }

// func (c *companyUsecase) GetCompanyInfo(ctx context.Context) (*models.Company, int, error) {
// 	company, err := c.companyRepo.Get(ctx)
// 	if err != nil {
// 		return nil, http.StatusNotFound, err
// 	}
// 	return company, http.StatusOK, err
// }

// func (c *companyUsecase) UpdateCompany(ctx context.Context, req request.CompanyRequest) (*models.Company, int, error) {
// 	company, err := c.companyRepo.Update(ctx & models.Company{
// 		Name:    req.Name,
// 		Balance: req.Balance,
// 	})

// 	if err != nil {
// 		return nil, http.StatusUnprocessableEntity, err
// 	}
// 	return company, http.StatusOK, nil
// }

// func (c *companyUsecase) CreateCompany(ctx context.Context, req request.CompanyRequest) (*models.Company, int, error) {
// 	company, err := c.companyRepo.Create(ctx & models.Company{
// 		Name:    req.Name,
// 		Balance: req.Balance,
// 	})

// 	if err != nil {
// 		return nil, http.StatusUnprocessableEntity, err
// 	}

// 	return company, http.StatusOK, nil
// }

// func (c *companyUsecase) TopupBalance(ctx context.Context, req *request.TopupCompanyBalance) (*models.Company, int, error) {
// 	company, err := c.companyRepo.AddBalance(ctx, req.Balance)
// 	if err != nil {
// 		return nil, http.StatusUnprocessableEntity, err
// 	}
// 	return company, http.StatusOK, nil
// }
