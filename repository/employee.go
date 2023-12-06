package repository

import (
	"finalproject_basisdata/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (c *EmployeeRepository) GetAllEmployee() ([]models.Employee, error) {
	var employees []models.Employee
	if err := c.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (c *EmployeeRepository) GetEmployeeById(id string) (*models.Employee, error) {
	var employee models.Employee
	if err := c.DB.Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (c *EmployeeRepository) CreateEmployeeData(employee *models.Employee) error {
	if err := c.DB.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

func (c *EmployeeRepository) UpdateEmployeeData(id string, req *models.EmployeeRequest) (*models.Employee, error) {
	var employee models.Employee
	if err := c.DB.Model(&employee).Where("id = ?", id).Updates(req).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (c *EmployeeRepository) DeleteEmployee(id string) error {
	var employee models.Employee
	if err := c.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
		return err
	}
	return nil

}
