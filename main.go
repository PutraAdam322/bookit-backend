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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := db.DBconnect()

	db.AutoMigrate(&model.User{}, &model.Booking{}, &model.BookingSlot{}, &model.Facility{})

	userRepo := repository.NewUserRepository(db)
	fctRepo := repository.NewFacilityRepository(db)
	bkngRepo := repository.NewBookingRepository(db)
	bkgsRepo := repository.NewBookingSlotRepository(db)

	jwtSvc := service.NewJWTService()
	userSvc := service.NewUserService(jwtSvc, userRepo)
	fctSvc := service.NewFacilityService(fctRepo)
	bkngSvc := service.NewBookingService(bkngRepo)
	bkgsSvc := service.NewBookingSlotService(bkgsRepo)

	userCtrl := controller.NewUserController(userSvc)
	fctCtrl := controller.NewFacilityController(fctSvc)
	bkngCtrl := controller.NewBookingController(bkngSvc, bkgsSvc)
	bkgSCtrl := controller.NewBookingSlotController(bkgsSvc)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	r.Use(middleware.CustomLogger())

	routes.UserRoutes(r, userCtrl, jwtSvc)
	routes.FacilityRoutes(r, fctCtrl, jwtSvc)
	routes.BookingRoutes(r, bkngCtrl, jwtSvc)
	routes.BookingSlotRoutes(r, bkgSCtrl, jwtSvc)
	routes.AdminRoutes(r, userCtrl, jwtSvc)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
