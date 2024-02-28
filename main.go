package main

import (
	"log"
	"os"

	"virtualminds/http-server/config"
	"virtualminds/http-server/database"
	"virtualminds/http-server/routers"
)

func main() {
	stdLogger := log.New(os.Stdout, "vm-sv:\t", log.Ldate|log.Ltime|log.Lshortfile)
	config.LoadConfig(stdLogger)

	dbRepo := &database.DatabaseRepository{}

	// Initialize database connection
	err := dbRepo.InitializeConnection(stdLogger)
	if err != nil {
		stdLogger.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer dbRepo.CloseAll()

	router := routers.SetupRoutes(stdLogger, dbRepo)
	router.Run(config.GetServicePort())
}
