package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID          int        `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	Employee_ID uuid.UUID  `gorm:"type:char(36);default:uuid()" json:"employee_id"`
	Name        *string    `json:"name"`
	Email       *string    `json:"email"`
	Phone       *string    `json:"phone"`
	Address     *string    `json:"address"`
	Private_Pin *string    `json:"private_pin"`
	Position_ID *string    `json:"position_id"`
	CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Private_Pin string `json:"private_pin"`
	Position_ID string `json:"position_id"`
}
