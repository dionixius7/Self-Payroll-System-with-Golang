package usecase

import (
	"errors"
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"log"
)

type EmployeeUsecase struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeUsecase(repo *repository.EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{Repo: repo}
}

func (c *EmployeeUsecase) GetAllEmployee() ([]models.Employee, error) {
	employees, err := c.Repo.GetAllEmployee()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (c *EmployeeUsecase) GetEmployeeById(id string) (*models.Employee, error) {
	employee, err := c.Repo.GetEmployeeById(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (c *EmployeeUsecase) CreateEmployeeData(req *models.EmployeeRequest) (*models.Employee, error) {
	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Address == "" || req.Private_Pin == "" || req.Position_ID == "" {
		return nil, errors.New("Silakan isi seluruh field")
	}
	employee := &models.Employee{
		Name:        &req.Name,
		Email:       &req.Email,
		Phone:       &req.Phone,
		Address:     &req.Address,
		Private_Pin: &req.Private_Pin,
		Position_ID: &req.Position_ID,
	}
	if err := c.Repo.CreateEmployeeData(employee); err != nil {
		log.Println("Tidak dapat membuat data karyawan: ", err)
		return nil, err
	}
	return employee, nil
}

func (c *EmployeeUsecase) UpdateEmployeeData(id string, req *models.EmployeeRequest) (*models.Employee, error) {
	employee, err := c.Repo.UpdateEmployeeData(id, req)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (c *EmployeeUsecase) DeleteEmployee(id string) error {
	err := c.Repo.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *EmployeeUsecase) WithdrawSalaryEmployee(id string, req *models.WithdrawRequest) error {
	return c.Repo.WithdrawSalaryEmployee(id, req)
}

// func (c *EmployeeUsecase) WithdrawSalaryEmployee(id string, req *models.WithdrawRequest) error {
// 	employee, err := c.Repo.WithdrawSalaryEmployee(&req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if employee.Private_Pin != req.Private_Pin {
// 		return errors.New("Pin yang anda masukkan salah")
// 	}
// 	position, err := c.Repo.WithdrawSalaryEmployee(&req)
// 	if err != nil {
// 		return nil, err
// 	}
// }
