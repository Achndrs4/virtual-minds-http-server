package main

import (
	"fmt"
	"log"
	"net/http"

	"virtualminds/http-server/pkg/database"
)

func main() {
	// Initialize the database connection
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	print(db)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
