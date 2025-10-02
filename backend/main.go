package main

import (
	"fmt"
	"log"

	"github.com/dylannguyennn/url-shortener/database"
	"github.com/dylannguyennn/url-shortener/router"
)

func main() {
	// Connect to DB
	if err := database.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Initialize router
	r := router.SetupRouter()

	fmt.Println("Server starting...")
	r.Run(":8080")
}
