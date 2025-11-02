package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func BookingSlotRoutes(router *gin.Engine, BookingSlotController *controller.BookingSlotController, jwtService middleware.JWTService) {
	bookingSlotRoutes := router.Group("/api/v1/booking_slots")
	{
		bookingSlotRoutes.POST("/create", middleware.Authenticate(jwtService), BookingSlotController.Create)
		bookingSlotRoutes.GET("/:id", BookingSlotController.GetByID)
		bookingSlotRoutes.GET("", middleware.Authenticate(jwtService), BookingSlotController.GetAll)
	}
}
