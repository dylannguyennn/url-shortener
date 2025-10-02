package main

import (
	"fmt"

	"github.com/dylannguyennn/url-shortener/router"
)

func main() {
	// Initialize router
	r := router.SetupRouter()

	fmt.Println("Server starting...")
	r.Run(":8080")
}
