package routes

import (
	"Gotest/controllers"
	"Gotest/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	api := router.Group("/user")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.GET("/:id", controllers.GetUserByID)
		api.GET("/getAll", controllers.GetAllUsers)
		api.GET("/info", middleware.AuthMiddleware(), controllers.GetUserInfoLogin)
		api.PUT("/update/:id", controllers.UpdateUserByID)
		api.PUT("/updateUser", middleware.AuthMiddleware(), controllers.UpdateUserByLogin)
		api.POST("/logout", controllers.LogoutUser)
	}
}
