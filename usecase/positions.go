package usecase

import (
	"errors"
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"log"
)

type PositionUsecase struct {
	Repo *repository.PositionRepository
}

func NewPositionUsecase(repo *repository.PositionRepository) *PositionUsecase {
	return &PositionUsecase{Repo: repo}
}

func (c *PositionUsecase) GetPositionById(id string) (*models.Position, error) {
	return c.Repo.GetPositionById(id)
}

func (c *PositionUsecase) GetAllPositionData() ([]*models.Position, error) {
	return c.Repo.GetAllPositionData()
}

func (c *PositionUsecase) CreatePosition(req *models.PositionReq) (*models.Position, error) {
	if req.NamePosition == nil || req.Salary == nil {
		return nil, errors.New("Isi seluruh kolom")
	}

	position := &models.Position{
		NamePosition: req.NamePosition,
		Salary:       req.Salary,
	}
	if err := c.Repo.CreatePosition(position); err != nil {
		log.Println("Gagal menambahkan data posisi baru: ", err)
		return nil, err
	}
	return position, nil
}

func (c *PositionUsecase) DeletePosition(id string) error {
	return c.Repo.DeletePosition(id)
}

func (c *PositionUsecase) UpdatePosition(id string, req *models.PositionReq) (*models.Position, error) {
	return c.Repo.UpdatePosition(id, req)
}

