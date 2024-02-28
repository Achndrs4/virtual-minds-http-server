package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig(logger *log.Logger, configFile string) {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(fmt.Errorf("fatal error config file: %s", err))
	}
}

func GetStandardLogger() *log.Logger {
	return log.New(os.Stdout, "SRVR-LOG:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func CreateDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.sql.user"),
		viper.GetString("database.sql.password"),
		viper.GetString("database.host"),
		viper.GetString("database.sql.port"),
		viper.GetString("database.name"))
	return dsn
}

func GetServicePort() string {
	return fmt.Sprintf(":%s", viper.GetString("server.port"))
}

func GetTrustedProxies() []string {
	return viper.GetStringSlice("server.trusted_proxies")
}

func GetGormSlowSQLSPeed() int {
	return viper.GetInt("server.logger.gorm.slow_sql_speed")
}

func GetConnectionBackoff() int {
	return viper.GetInt("server.connection_backoff")
}

func GetMaxConnectionRetries() int {
	return viper.GetInt("server.max_connection_retries")
}
