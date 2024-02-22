package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	// Initialize Viper for reading configuration
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	dbConfig := viper.GetStringMapString("database")

	// Connect to the database with connection pooling
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig["host"], dbConfig["user"], dbConfig["password"], dbConfig["dbname"], dbConfig["port"])
	print(dsn)
	return nil, nil
}

func getViperConfiguration(string, string) {

}
