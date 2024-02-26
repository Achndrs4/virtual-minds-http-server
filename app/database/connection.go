package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"virtualminds/http-server/models"
)

var readerDB *gorm.DB
var writerDB *gorm.DB

func ConnectToWriterDB(logger *log.Logger, dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Could not connect to Writer Database:\t%s", err)
	}

	// migrate ORM settings
	db.AutoMigrate(&models.Customer{}, &models.HourlyStat{}, &models.IPBlacklist{}, &models.UABlacklist{})

	logger.Print("Connected to the Master Database ...\n")
	writerDB = db
	return db
}

func ConnectToReaderDB(logger *log.Logger, dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Could not connect to Writer Database:\t%s", err)
	}

	// migrate ORM settings
	db.AutoMigrate(&models.Customer{}, &models.HourlyStat{}, &models.IPBlacklist{}, &models.UABlacklist{})
	logger.Print("Connected to the Replica Database ...\n")
	readerDB = db
	return db
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
