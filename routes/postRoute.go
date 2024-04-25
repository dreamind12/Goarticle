package routes

import (
	"Gotest/controllers"
	"Gotest/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	api := router.Group("/post")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/create", controllers.CreatePost)
		api.GET("/:id", controllers.GetPostByID)
		api.GET("/page/:limit/:offset", controllers.GetPostsWithPaging)
		api.GET("/getAll", controllers.GetAllPosts)
		api.PUT("/:id", controllers.UpdatePostByID)
		api.DELETE("/:id", controllers.DeletePostByID)
	}
}
