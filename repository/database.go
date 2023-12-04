package repository

import (
	"finalproject_basisdata/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/finalproject_basisdata?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(
		&models.Company{},
		&models.Transaction{},
		&models.Employee{},
		&models.Admin{},
		&models.Position{},
	)
	DB = database

	db, err := database.DB()
	if err != nil {
		panic("connect database :" + err.Error())
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}

// func DisconnectDatabase() string {
// 	if DB != nil {
// 		sqlDB, err := DB.DB()
// 		if err != nil {
// 			panic(err)
// 		}
// 		sqlDB.Close()
// 		return "Sukses memutuskan koneksi dari database."
// 	}
// 	return "Tidak ada koneksi yang perlu diputus."
// }
