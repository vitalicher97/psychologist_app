package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vitalicher97/psychologist_app/internal/app/api"
	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDBConnection(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	defer db.GetConnection().Close()

	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/image/:imageName", func(c *gin.Context) {
		imageName := c.Param("imageName")
		imagePath := "./static/images/profile/" + imageName

		// Serve the image file
		c.File(imagePath)
	})

	api.SetupRoutes(r)

	log.Println("Server is running...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
