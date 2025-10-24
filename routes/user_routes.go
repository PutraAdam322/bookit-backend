package routes

import (
	"bookit.com/controller"
	"bookit.com/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/api/v1/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
	}
}
