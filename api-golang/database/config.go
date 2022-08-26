package database

import (
	"log"

	"github.com/devmontini/wiselink-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=montini password=123123 dbname=montini_db port=5433"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Database connection successful")

	db.AutoMigrate(&models.Users{}, &models.Events{}, &models.UsersEvents{})

	db.SetupJoinTable(&models.Users{}, "Eventss", &models.UsersEvents{})
	DB = db
}
