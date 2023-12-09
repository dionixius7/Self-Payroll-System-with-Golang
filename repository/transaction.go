package repository

import (
	"finalproject_basisdata/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (c *TransactionRepository) GetInfoTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := c.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
