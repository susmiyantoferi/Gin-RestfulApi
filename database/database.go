package database

import (
	"RestApi-Gin/helper"
	"RestApi-Gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDb() *gorm.DB {
	conn := "host=localhost user=postgres password=Terserah123 dbname=gin_gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	helper.PanicError(err)
	erru := db.AutoMigrate(&model.Mahasiswa{})
	helper.PanicError(erru)

	return db
}
