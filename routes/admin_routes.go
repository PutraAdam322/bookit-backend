package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, userController *controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/admin")
	{
		userRoutes.POST("/login", userController.AdminLogin)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
	}
}
