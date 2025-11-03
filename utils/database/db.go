package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBconnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST_LOCAL")
	port := os.Getenv("DB_PORT_LOCAL")
	dbName := os.Getenv("DB_DATABASE_LOCAL")
	username := os.Getenv("DB_USERNAME_LOCAL")
	password := os.Getenv("DB_PASSWORD_LOCAL")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error initiating database : ", err)
		os.Exit(1)
	}

	return db
}
