package models

import "time"

// const (
// 	TopupTransactionType string = "debit"
// 	TopupTransactionNote string = "topup saldo perusahaan"
// )

type Transaction struct {
	ID        int        `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	Type      string     `gorm:"type:varchar(36)" json:"type"`
	Amount    *int       `gorm:"type:int" json:"amount"`
	Note      string     `gorm:"type:varchar(36)" json:"note"`
	CreatedAt time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type TransactionRequest struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Note   string `json:"note"`
}
