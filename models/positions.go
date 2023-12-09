package models

import "time"

type Position struct {
	ID           int        `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	NamePosition *string    `gorm:"type:varchar(36)" json:"name_position"`
	Salary       *int       `gorm:"type:int" json:"salary"`
	CreatedAt    time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type PositionReq struct {
	NamePosition *string `json:"name_position" validate:"required"`
	Salary       *int    `json:"salary" validate:"required"`
}
