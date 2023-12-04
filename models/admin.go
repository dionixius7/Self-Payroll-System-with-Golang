package models

import "github.com/google/uuid"

type Admin struct {
	ID       int       `gorm:"AUTO_INCREMENT;primaryKey" json:"id"`
	AdminID  uuid.UUID `gorm:"type:char(36);default:uuid()" json:"admin_id"`
	Email    string    `gorm:"type:varchar(300);unique;" json:"email"`
	Name     string    `gorm:"type:varchar(300);" json:"name"`
	Password string    `gorm:"type:text;" json:"password"`
}

type ReqSignUp struct {
	Name     string `gorm:"type:varchar(300);" json:"name"`
	Email    string `gorm:"type:varchar(300);unique;" json:"email"`
	Password string `gorm:"type:text;" json:"password"`
}

type LoginPayload struct {
	Email    string `gorm:"type:varchar(300);unique" json:"email" validate:"required"`
	Password string `gorm:"type:text" json:"password" validate:"required"`
}
