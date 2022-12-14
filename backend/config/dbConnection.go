package config

import (
	"fmt"
	"jwt-middleware/backend/Models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dbhost := os.Getenv("POSTGRE_HOST")
	dbuser := os.Getenv("POSTGRE_USER")
	dbpassword := os.Getenv("POSTGRE_PASSWORD")
	dbname := os.Getenv("POSTGRE_DBNAME")
	dbport := os.Getenv("POSTGRE_DBPORT")
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", dbhost, dbuser, dbpassword, dbname, dbport)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)

	}
	DB = db

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {

	connection.Debug().AutoMigrate(
		&Models.User{},
	)
	fmt.Println("Tablolar Oluşturuldu")

}
