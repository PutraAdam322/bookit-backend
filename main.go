package main

import (
	"log"

	"bookit.com/controller"
	"bookit.com/middleware"
	"bookit.com/model"
	"bookit.com/repository"
	"bookit.com/routes"
	"bookit.com/service"
	db "bookit.com/utils/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := db.DBconnect()

	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)

	jwtSvc := service.NewJWTService()
	userSvc := service.NewUserService(jwtSvc, userRepo)

	userCtrl := controller.NewUserController(userSvc)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	routes.UserRoutes(r, userCtrl, jwtSvc)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
