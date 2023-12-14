package controllers

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	Usecase *usecase.EmployeeUsecase
}

func NewEmployeeController(usecase *usecase.EmployeeUsecase) *EmployeeController {
	return &EmployeeController{Usecase: usecase}
}

// func (c *EmployeeController) GetAllEmployee(ctx *fiber.Ctx) error {
// 	employees, err := c.Usecase.GetAllEmployee()

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Terdapat kesalahan dalam database",
// 		})
// 	}
// 	fmt.Println("Tipe data dari employees:", reflect.TypeOf(employees))

// 	// return ctx.JSON(fiber.Map{
// 	// 	"status": fiber.StatusOK,
// 	// 	"data":   employees,
// 	// })
// 	return ctx.JSON(employees)

// }
func (c *EmployeeController) GetAllEmployee(ctx *fiber.Ctx) error {
	employees, err := c.Usecase.GetAllEmployee()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terdapat kesalahan dalam database",
		})
	}
	return ctx.JSON(employees)
}

func (c *EmployeeController) GetEmployeeById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	employee, err := c.Usecase.GetEmployeeById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}
	// return ctx.JSON(fiber.Map{
	// 	"status": fiber.StatusOK,
	// 	"data":   employee,
	// })
	return ctx.JSON(employee)
}

func (c *EmployeeController) CreateEmployeeData(ctx *fiber.Ctx) error {
	var req *models.EmployeeRequest
	log.Println("Recieve req body: ", string(ctx.Body()))
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Address == "" || req.Private_Pin == "" || req.Position_ID == "" {
		log.Println("Harap isi seluruh kolom")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Harap isi seluruh Kolom",
		})
	}
	employee, err := c.Usecase.CreateEmployeeData(req)
	if err != nil {
		log.Println("Gagal membuat data karyawan: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat data karyawan",
		})
	}
	return ctx.JSON(employee)
	// return ctx.JSON(fiber.Map{
	// 	"message": fiber.StatusOK,
	// 	"data":    employee,
	// })
}

func (c *EmployeeController) UpdateEmployeeData(ctx *fiber.Ctx) error {
	var req *models.EmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	id := ctx.Params("id")
	employee, err := c.Usecase.UpdateEmployeeData(id, req)
	if err != nil {
		log.Println("Gagal memuat data karyawan : ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Karyawan tidak ditemukan",
		})
	}
	return ctx.JSON(employee)
	// return ctx.JSON(fiber.Map{
	// 	"message": "Berhasil mengubah data karyawan",
	// 	"data":    employee,
	// })
}

func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.Usecase.DeleteEmployee(id)
	if err != nil {
		log.Println("Gagal memuat data karyawan: ", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data karyawan tidak ditemukan",
		})
	}
	return ctx.JSON("Sukses menghapus data karyawan")
	// return ctx.JSON(fiber.Map{
	// 	"message": "Sukses menghapus data karyawan",
	// })
}

func (c *EmployeeController) WithdrawSalaryEmployee(ctx *fiber.Ctx) error {
	var req models.WithdrawRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Invalid req body: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	id := ctx.Params("id")
	err := c.Usecase.WithdrawSalaryEmployee(id, &req)
	if err != nil {
		log.Println("Err: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON("Berhasil mencairkan gaji bulanan Anda")
	// return ctx.JSON(fiber.Map{
	// 	"message": fiber.StatusOK,
	// 	"data":    "Berhasil mencairkan gaji bulanan Anda",
	// })
}

//cek
