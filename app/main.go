package main

import (
	"log"
	"os"

	"virtualminds/http-server/config"
	"virtualminds/http-server/database"
	"virtualminds/http-server/routers"
)

func main() {
	mainLogger := log.New(os.Stdout, "vm-sv:\t", log.Ldate|log.Ltime|log.Lshortfile)
	config.Initialize(mainLogger)

	database.ConnectToWriterDB(mainLogger, config.CreateWriterDSN())
	database.ConnectToReaderDB(mainLogger, config.CreateReaderDSN())
	defer database.Close()

	router := routers.SetupRoutes(mainLogger)
	router.Run(config.GetServicePort())
}
