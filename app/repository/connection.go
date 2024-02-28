package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"virtualminds/http-server/config"
	"virtualminds/http-server/models"
)

var readerDB *gorm.DB
var writerDB *gorm.DB
var gormLogger logger.Interface

func InitializeConnection(std_logger *log.Logger) {
	configureGormLogger(std_logger)
	var err error
	writerDB, err = connectToDB(config.CreateDSN())
	if err != nil {
		std_logger.Fatalf("Could not connect to master writer database: %s", err.Error())
	}
	readerDB, err = connectToDB(config.CreateDSN())
	if err != nil {
		std_logger.Fatalf("Could not connect to replica reader database: %s", err.Error())
	}
}
func connectToDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		return nil, err
	}

	// migrate ORM settings
	err = db.AutoMigrate(&models.Customer{}, &models.HourlyStat{}, &models.IPBlacklist{}, &models.UABlacklist{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func configureGormLogger(std_logger *log.Logger) {
	gormLogger = logger.New(
		std_logger,
		logger.Config{
			// seconds for an SQL query to be considered "slow" - change as queries get more heavy
			SlowThreshold: time.Duration(config.GetGormSlowSQLSPeed()) * time.Second,
			// we don't want to flood the logs, and anyways we store good and bad requests
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}

func GetReaderDB() *gorm.DB {
	return readerDB
}

func GetWriterDB() *gorm.DB {
	return writerDB
}

func Close() error {
	reader, err := readerDB.DB()
	if err == nil {
		reader.Close()
	} else {
		return err
	}

	writer, err := writerDB.DB()
	if err == nil {
		writer.Close()
	} else {
		return err
	}
	return nil
}
