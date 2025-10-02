package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dylannguyennn/url-shortener/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	// Connect to DB
	if err := database.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Gin router with default middleware
	r := gin.Default()

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET
	// Creates shortened URL
	r.POST("/shorten", func(c *gin.Context) {
		// Struct for parsing and validating URL
		var req struct {
			URL string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Generates shortID by trimming UUID and taking first 8 chars
		shortID := uuid.New().String()[:8]

		// Inserts shortID and URL into short_id and original_url in DB
		_, err := database.DB.Exec(
			"INSERT INTO urls (short_id, original_url) VALUES (?, ?)",
			shortID, req.URL,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into DB"})
			return
		}

		// Returns shortened URL
		c.JSON(http.StatusOK, gin.H{
			"short_url": "https://localhost:8080" + shortID,
			"original":  req.URL,
		})
	})

	// GET
	// URL redirection
	r.GET("/:shortID", func(c *gin.Context) {
		shortID := c.Param("shortID")

		// Looks up original_url using short_id
		row := database.DB.QueryRow(
			"SELECT original_url FROM urls WHERE short_id = ?",
			shortID,
		)

		var original string
		err := row.Scan(&original)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}

		// Redirects to original URL
		c.Redirect(http.StatusMovedPermanently, original)
	})

	fmt.Println("Server starting...")
	r.Run(":8080")
}
