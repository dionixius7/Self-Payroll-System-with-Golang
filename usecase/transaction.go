package usecase

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
)

type TransactionUsecase struct {
	Repo *repository.TransactionRepository
}

func NewTransactionUsecase(repo *repository.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{Repo: repo}
}

func (c *TransactionUsecase) GetInfoTransaction() ([]models.Transaction, error) {
	return c.Repo.GetInfoTransaction()
}
