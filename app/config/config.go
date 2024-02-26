package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Initialize(logger *log.Logger) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Errorf("fatal error config file: %s", err))
	}
}

func PrintConfig() {
	allSettings := viper.AllSettings()
	for key, value := range allSettings {
		if strings.Contains(key, "password") {
			fmt.Printf("%s: %v\n", key, "****")
		} else {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

func CreateWriterDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.writer.user"),
		viper.GetString("database.writer.password"),
		viper.GetString("database.host"),
		viper.GetString("database.writer.port"),
		viper.GetString("database.name"))
	return dsn
}

func CreateReaderDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.reader.user"),
		viper.GetString("database.reader.password"),
		viper.GetString("database.host"),
		viper.GetString("database.reader.port"),
		viper.GetString("database.name"))
	return dsn
}

func GetServicePort() string {
	return fmt.Sprintf(":%s", viper.GetString("app.port"))
}
