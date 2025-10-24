package main

import (
	"fmt"

	"bookit.com/model"
	db "bookit.com/utils/database"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(p string) (string, error) {
	pwd := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MaxCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func main() {
	godotenv.Load()
	db := db.DBconnect()
	fmt.Println("initiating...")

	db.AutoMigrate(&model.User{})
	fmt.Println("database is migrated")

	pwd := "12345678"
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("cannot hash password")
		return
	}
	fmt.Println("password is hashed successfully")

	var users []model.User
	users = append(users, model.User{
		Name:         "Admin1",
		PasswordHash: string(pwdHash),
		Email:        "admin1@gmail.com",
		IsAdmin:      true,
	})
	users = append(users, model.User{
		Name:         "User1",
		PasswordHash: string(pwdHash),
		Email:        "user1@gmail.com",
		IsAdmin:      false,
	})
	fmt.Println("data is inserted...")

	db.Create(&users)

}
