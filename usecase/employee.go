package usecase

import (
	"errors"
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"
	"log"
)

type EmployeeUsecase struct {
	Repo *repository.EmployeeRepository
	Company *repository.CompanyRepository
	//Employee *repository.UpdateEmployeeData
}
func NewEmployeeUsecase(repo *repository.EmployeeRepository, companyRepo *repository.CompanyRepository) *EmployeeUsecase {
	return &EmployeeUsecase{
		Repo: repo,
		Company: companyRepo,
		//Employee: employeeRepo,
	}
}
// func NewEmployeeUsecase(repo *repository.EmployeeRepository, companyRepo *repository.CompanyRepository) *EmployeeUsecase {
// 	return &EmployeeUsecase{
// 		Repo: repo,
// 		Company: companyRepo,
// 		//Employee: employeeRepo,
// 	}
// }

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
	// Convert EmployeeRequest to Employee
	employee := &models.Employee{
		Name:        &req.Name,
		Email:       &req.Email,
		Phone:       &req.Phone,
		Address:     &req.Address,
		Private_Pin: &req.Private_Pin,
		Position_ID: &req.Position_ID,
		//Is_Paid:     &req.Is_Paid,
	}

	// Call repository method with the converted employee
	updatedEmployee, err := c.Repo.UpdateEmployeeData(id, employee)
	if err != nil {
		return nil, err
	}

	return updatedEmployee, nil
}

func (c *EmployeeUsecase) DeleteEmployee(id string) error {
	err := c.Repo.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *EmployeeUsecase) WithdrawSalaryEmployee(id string, req *models.WithdrawRequest) error {
	employee, err := c.Repo.GetEmployeeById(id)
	if err != nil {
		return errors.New("Data karyawan tidak ditemukan")
	}

	if *employee.Private_Pin != req.Private_Pin {
		return errors.New("Pin yang anda masukkan salah")
	}

	position, err := c.Repo.Position.GetPositionById(id)
	if err != nil {
		return errors.New("Data posisi tidak ditemukan")
	}

	if *employee.Is_Paid {
		return errors.New("Anda sudah mencairkan gaji pada bulan ini")
	}

	salary := position.Salary
	if salary == nil {
		return errors.New("Gaji karyawan belum ditetapkan")
	}

	employee.Is_Paid = new(bool)
	*employee.Is_Paid = true
	_, err2 := c.Repo.UpdateEmployeeData(id, employee)
	if err2 != nil {
		return errors.New("Gagal menyimpan perubahan status gaji")
	}

	company, err := c.Company.GetCompanyInfo("1")
	if err != nil {
		return err
	}

	err = c.Company.UpdateCompanyBalanceAfterWithdraw(*company, *salary)
	if err != nil {
		return err
	}
	return nil
}
