package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func FacilityRoutes(router *gin.Engine, facilityController *controller.FacilityController, jwtService middleware.JWTService) {
	facilityRoutes := router.Group("/api/v1/facilities")
	{
		facilityRoutes.GET("/:id", middleware.Authenticate(jwtService), facilityController.GetByID)
		facilityRoutes.POST("/create", middleware.Authenticate(jwtService), facilityController.Create)
		facilityRoutes.GET("", facilityController.GetAll)
	}
}
