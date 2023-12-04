package models

import "time"

type Transaction struct {
	ID        int        `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	Type      *string    `gorm:"type:varchar(36)" json:"type"`
	Amount    *int       `gorm:"type:int" json:"amount"`
	Note      *string    `gorm:"type:varchar(36)" json:"note"`
	CreatedAt time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
