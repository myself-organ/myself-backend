package main

import (
	"fmt"
	"log"
	"myself-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dataSourceName := "myself-backend.db"
	repo, err := repository.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	fmt.Printf("Repository initialized: %v\n", repo)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("cv")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Save the uploaded file to a specific destination.
		err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "success"})
	})

	router.Run(":8081")
}
