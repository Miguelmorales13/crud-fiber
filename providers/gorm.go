package providers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	Database *gorm.DB
)

func InitDbPostgres() {
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSMODE"),
	)
	var err error = nil
	Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("[ERROR_DATABASE_CONNECTION][", err.Error(), "]")
	}
	err = Database.AutoMigrate()
	if err != nil {
		log.Fatal(err.Error())
	}
}
