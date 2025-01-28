package main

import (
	"fmt"
	"log"
	"myself-backend/internal/repository"
	"strconv"

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

	router.GET("/cv/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		cv, err := repo.FindByID(intID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, cv)
	})

	router.POST("/cv", func(c *gin.Context) {
		var cv repository.CV
		if err := c.BindJSON(&cv); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := repo.Save(cv); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "success"})
	})

	router.GET("/cvs", func(c *gin.Context) {
		cvs, err := repo.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, cvs)
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
