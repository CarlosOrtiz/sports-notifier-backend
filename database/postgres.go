package database

import (
	"log"

	"github.com/CarlosOrtiz/sports-notifier-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionPostgres() {
	var error error

	DB, error = gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected Postgres")
	}
}
