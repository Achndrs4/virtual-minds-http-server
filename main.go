package main

import (
	"virtualminds/http-server/config"
	"virtualminds/http-server/database"
	"virtualminds/http-server/routers"
)

func main() {
	// Initialize configuration
	stdLogger := config.GetStandardLogger()
	config.LoadConfig(stdLogger, "config.yaml")

	// Initialize database connection
	dbRepo := &database.DatabaseRepository{}
	err := dbRepo.InitializeConnection(stdLogger)
	if err != nil {
		stdLogger.Fatalf("Failed to initialize database connection: %v", err)
	}
	// close database connection at the end
	defer dbRepo.CloseAll()

	// setup routes
	router := routers.SetupRoutes(stdLogger, dbRepo)

	// serve routes
	router.Run(config.GetServicePort())
}
