package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine, BookingController *controller.BookingController, jwtService middleware.JWTService) {
	bookingRoutes := router.Group("/api/v1/bookings")
	{
		bookingRoutes.POST("/create", middleware.Authenticate(jwtService), BookingController.Create)
		bookingRoutes.GET("/mybook", middleware.Authenticate(jwtService), BookingController.GetByUserID)
		bookingRoutes.GET("/admin/bookings", middleware.Authenticate(jwtService), BookingController.GetAll)
		bookingRoutes.PATCH(("/:id/cancel"), middleware.Authenticate(jwtService), BookingController.Cancel)
	}
}
