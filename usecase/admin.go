package usecase

// import (
// 	"errors"
// 	"finalproject_basisdata/models"
// 	"finalproject_basisdata/repository"
// 	"log"
// 	"regexp"

// 	"golang.org/x/crypto/bcrypt"
// )

// type AdminUsecase struct {
// 	Repo *repository.AdminRepository
// }

// func NewAdminUsecase(repo *repository.AdminRepository) *AdminUsecase {
// 	return &AdminUsecase{Repo: repo}
// }

// // koreksi nih
// func (c *AdminUsecase) SignUpAdmin(req models.ReqSignUp) (*models.Admin, error) {
// 	if req.Email == "" || req.Name == "" || req.Password == "" {
// 		return nil, errors.New("Form tidak lengkap")
// 	}
// 	admin := &models.Admin{
// 		Email:    req.Email,
// 		Name:     req.Name,
// 		Password: req.Password,
// 	}
// 	err := c.Repo.SignUpAdmin(admin)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return admin, nil
// }

// func (c *AdminUsecase) CekEmail(req *models.ReqSignUp) error {
// 	admin, err := c.Repo.CekEmail(req.Email)
// 	if err != nil {
// 		log.Println("Err: ", err)
// 		return err
// 	}
// 	if admin != nil {
// 		return errors.New("Email sudah terdaftar: " + req.Email)
// 	}
// 	return nil
// }

// func (c *AdminUsecase) CekName(name string) error {
// 	admin, err := c.Repo.CekName(name)
// 	if err != nil {
// 		log.Println("Nama sudah terdaftar: ", err)
// 		return err
// 	}
// 	if admin != nil {
// 		return errors.New("Nama sudah terdaftar: " + name)
// 	}
// 	return nil
// }

// func (c *AdminUsecase) CekPassword(password string) error {
// 	const (
// 		minLength      = 8
// 		minUppercase   = 1
// 		minDigits      = 1
// 		minSpecialChar = 1
// 	)

// 	if len(password) < minLength {
// 		return errors.New("Password kurang panjang!")
// 	}
// 	uppercaseCount := len(regexp.MustCompile(`[^A-Z]`).ReplaceAllString(password, ""))
// 	if uppercaseCount < minUppercase {
// 		return errors.New("Masukkan minimal 1 huruf kapital pada password!")
// 	}
// 	digitCount := len(regexp.MustCompile(`[^0-9]`).ReplaceAllString(password, ""))
// 	if digitCount < minDigits {
// 		return errors.New("Masukkan minimal 1 angka pada password!")
// 	}
// 	specialCharCount := len(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(password, ""))
// 	if specialCharCount < minSpecialChar {
// 		return errors.New("Masukkan minimal 1 karakter spesial pada password!")
// 	}
// 	return nil
// }

// func (c *AdminUsecase) HashPassword(password string) (string, error) {
// 	if len(password) < 8 {
// 		return "", bcrypt.ErrHashTooShort
// 	}
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(hashedPassword), err
// }

// func (c *AdminUsecase) CheckPasswordHash(password, hash string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// }
