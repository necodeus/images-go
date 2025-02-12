package main

import (
	"images-go/database"
	"images-go/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	uploadsDir := os.Getenv("UPLOADS_DIR")
	os.MkdirAll(uploadsDir, os.ModePerm)

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.InitDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer database.Close()

	// router
	router := gin.Default()
	router.GET("/:resource_id", handlers.GetImage)
	router.GET("/:resource_id/:resolution", handlers.GetImage)

	router.Run(":8080")
}
