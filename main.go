package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"Gotest/controllers"
	"Gotest/database"
)

func main() {
	database.InitDB()

	// Membuat tabel artikel
	// database.CreateArticleTable()

	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	})

	// Menambahkan rute dan handler
	api := router.Group("/article")
	{
		api.POST("/", controllers.CreatePost)
		api.GET("/:id", controllers.GetPostByID)
		api.GET("/page/:limit/:offset", controllers.GetPostsWithPaging)
		api.PUT("/:id", controllers.UpdatePostByID)
		api.DELETE("/:id", controllers.DeletePostByID)
	}

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Response success")
		c.String(http.StatusOK, "Response Success!")
	})

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
