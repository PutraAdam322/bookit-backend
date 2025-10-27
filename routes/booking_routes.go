package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine, BookingController *controller.BookingController, jwtService middleware.JWTService) {
	bookingRoutes := router.Group("/api/v1/bookings")
	{
		bookingRoutes.GET("/:id", BookingController.GetAll)
		bookingRoutes.POST("/create", middleware.Authenticate(jwtService), BookingController.Create)
		bookingRoutes.GET("", middleware.Authenticate(jwtService), BookingController.GetAll)
	}
}
