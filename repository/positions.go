package repository

import (
	"finalproject_basisdata/models"

	"gorm.io/gorm"
)

type PositionRepository struct {
	DB *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{DB: db}
}
func (c *PositionRepository) CreatePosition(position *models.Position) error {
	if err := c.DB.Create(&position).Error; err != nil {
		return err
	}

	return nil
}


func (c *PositionRepository) DeletePosition(id string) error {
	result := c.DB.Where("id = ?", id).Delete(&models.Position{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *PositionRepository) UpdatePosition(id string, req *models.PositionReq) (*models.Position, error) {
	position := &models.Position{
		NamePosition: req.NamePosition,
		Salary:       req.Salary,
	}

	if err := c.DB.Model(&models.Position{}).Where("id = ?", id).Updates(position).Error; err != nil {
		return nil, err
	}

	return position, nil
}

func (c *PositionRepository) GetPositionById(id string) (*models.Position, error) {
	var position models.Position
	if err := c.DB.First(&position, id).Error; err != nil {
		return nil, err
	}

	return &position, nil
}
