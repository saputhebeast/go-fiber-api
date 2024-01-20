package database

import (
	"go-fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectionDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!\n", err.Error())
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("Failed to migrate models!\n", err.Error())
	}

	Database = DbInstance{Db: db}
}
