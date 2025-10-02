package main

import (
	"fmt"
	"log"

	"github.com/dylannguyennn/url-shortener/database"
)

func main() {
	// Connect to DB
	if err := database.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize router
	r := router.setupRouter()

	fmt.Println("Server starting...")
	r.Run(":8080")
}
