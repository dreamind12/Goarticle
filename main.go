package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"Gotest/database"
	"Gotest/routes"
)

func main() {
	database.InitDB()

	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router.Use(cors.Default())

	router.Use(func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	})

	// Routes
	routes.PostRoute(router)
	routes.UserRoutes(router)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("Response success")
		c.String(http.StatusOK, "Response Success!")
	})

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
