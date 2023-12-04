package models

import (
	"time"
)

type Company struct {
	ID        int        `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(300);" json:"name"`
	Balance   *int       `gorm:"type:int" json:"balance"`
	CreatedAt time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CompanyRequest struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type TopupCompanyBalance struct {
	Balance int `json:"balance;" validate:"required"`
}

// CompanyRepository interface {
// 	Get(ctx context.Context) (*Company, error)
// 	Create(ctx context.Context, Company *Company) (*Company, error)
// 	Update(ctx context.Context, Company *Company) (*Company, error)
// 	Delete(ctx context.Context, balance int) (*Company, error)
// 	DebitBalance(ctx context.Context, amount int, note string) error
// }

// CompanyUsecase interface {
// 	GetCompanyInfo(ctx context.Context) (*Company, int, error)
// 	CreateCompany(ctx context.Context, req request.CompanyRequest) (*Company, int, error)
// 	UpdateCompany(ctx context.Context, req request.CompanyRequest) (*Company, int, error)
// 	TopupBalance(ctx context.Context, req *request.TopupCompanyBalance) (*Company, int, error)
// }
