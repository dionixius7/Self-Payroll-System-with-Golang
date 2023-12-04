package models

import (
	"time"
)

type Company struct {
	ID      int    `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(300);" json:"name"`
	Balance *int   `gorm:"type:int" json:"balance"`
	// Note      string     `gorm:"type:varchar(300);" json:"note"`
	CreatedAt time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CompanyRequest struct {
	Name    string `json:"name" validate:"required"`
	Balance int    `json:"balance" validate:"required"`
}

type TopupCompanyBalance struct {
	Balance int `json:"balance;" validate:"required"`
}
