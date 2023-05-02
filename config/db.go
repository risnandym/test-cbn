package config

import (
	"os"
	"test-case/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {

	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	database := os.Getenv("DATABASE_NAME")
	ssl := os.Getenv("SSL_MODE")

	// production
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=" + ssl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Product{}, &models.Image{})

	return db

}
