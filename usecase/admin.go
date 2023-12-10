package usecase

import (
	"finalproject_basisdata/models"
	"finalproject_basisdata/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	Repo *repository.AdminRepository
}

func NewAdminUsecase(repo *repository.AdminRepository) *AdminUsecase {
	return &AdminUsecase{Repo: repo}
}

func (c *AdminUsecase) SignUpAdmin(req *models.ReqSignUp) error {
	err := c.Repo.CheckEmailExists(req.Email)
	if err != nil {
		return err
	}

	err = c.Repo.CheckNameExists(req.Name)
	if err != nil {
		return err
	}

	if err := c.Repo.CheckPassword(req.Email, req.Password); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return err
	}

	newAdmin := &models.Admin{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
	}

	if err := c.Repo.CreateAdmin(newAdmin); err != nil {
		return err
	}

	return nil
}

func (c *AdminUsecase) LoginAdmin(req *models.LoginPayload) error {
	err := c.Repo.CheckEmailLogin(req.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := c.Repo.GetHashedPassword(req.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return err
	}

	return nil
}
