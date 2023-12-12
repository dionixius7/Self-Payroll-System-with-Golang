package repository

import (
	//"errors"
	"finalproject_basisdata/models"

	//"finalproject_basisdata/repository/position"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB       *gorm.DB
	Position *PositionRepository
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		DB:       db,
		Position: NewPositionRepository(db),
	}
}

//Position: position.NewPositionRepository(db),

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

// func (c *EmployeeRepository) WithdrawSalaryEmployee(id string, req *models.WithdrawRequest) error {
// 	employee, err := c.GetEmployeeById(id)
// 	if err != nil {
// 		return errors.New("Data karyawan tidak ditemukan")
// 	}

// 	if *employee.Private_Pin != req.Private_Pin {
// 		return errors.New("Pin yang anda masukkan salah")
// 	}

// 	position, err := c.Position.GetPositionById(id) // Corrected method name
// 	if err != nil {
// 		return errors.New("Data posisi tidak ditemukan")
// 	}

// 	salary := position.Salary
// 	if salary == nil {
// 		return errors.New("Gaji karyawan belum ditetapkan")
// 	}

// 	if err := c.DB.Save(&position).Error; err != nil {
// 		return errors.New("Gagal menarik gaji")
// 	}
	
// 	return nil
// }
