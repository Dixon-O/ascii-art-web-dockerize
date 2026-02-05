package main

import (
	"ascii-art/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize templates
	err := handlers.InitTemplates()
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	// Register routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	// Start server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
