package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"virtualminds/http-server/config"
	"virtualminds/http-server/models"
	"virtualminds/http-server/utils"
)

type DatabaseInterface interface {
	InitializeConnection(std_logger *log.Logger) error
	UpdateStatsCount(request *models.CustomerRequest, isValid bool) error
	GetCustomer(customerID uint) (*models.Customer, error)
	GetCustomerTotal(datetime time.Time, customer int) (*models.CustomerStats, error)
	GetDailyTotal(datetime time.Time) (int64, error)
	IsIPBanned(ip int) (bool, error)
	IsUserAgentBanned(userAgent string) (bool, error)
	CloseAll() error
}

type DatabaseRepository struct {
	DB         *gorm.DB
	gormLogger logger.Interface
}

func (r *DatabaseRepository) InitializeConnection(std_logger *log.Logger) error {
	r.configureGormLogger(std_logger)

	var err error
	r.DB, err = r.connectToDB(config.CreateDSN())
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) connectToDB(dsn string) (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: r.gormLogger})
	if err != nil {
		return nil, err
	}

	// migrate ORM settings
	err = database.AutoMigrate(&models.Customer{}, &models.HourlyStat{}, &models.IPBlacklist{}, &models.UABlacklist{})
	if err != nil {
		return nil, err
	}
	return database, nil
}

func (r *DatabaseRepository) configureGormLogger(std_logger *log.Logger) {
	r.gormLogger = logger.New(
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

// CloseAll closes all database connections
func (r *DatabaseRepository) CloseAll() error {
	conn, err := r.DB.DB()
	if err == nil {
		conn.Close()
		return nil
	} else {
		return err
	}
}

// UpdateStatsCount updates the hourly stats count
func (r *DatabaseRepository) UpdateStatsCount(request *models.CustomerRequest, isValid bool) error {
	database := r.DB
	result := &models.HourlyStat{}

	// Round down the request timestamp to the nearest hour
	hour := utils.RoundDownToHour(request.Timestamp)

	// If hourly stats already exist for the given hour, retrieve them. Otherwise, create new stats.
	if err := database.FirstOrCreate(result, &models.HourlyStat{CustomerID: request.CustomerID, Time: hour}).Error; err != nil {
		return err
	}

	// Increment the request count
	result.RequestCount++

	// Increment the invalid count if necessary
	if !isValid {
		result.InvalidCount++
	}

	// Save to the database
	if err := database.Save(result).Error; err != nil {
		return err
	}
	return nil
}

// GetCustomer retrieves a customer by ID
func (r *DatabaseRepository) GetCustomer(customerID uint) (*models.Customer, error) {
	database := r.DB
	var customer models.Customer
	result := database.First(&customer, customerID)
	if result.Error == nil {
		return &customer, nil
	}
	return nil, result.Error
}

// IsIPBanned checks if an IP is banned
func (r *DatabaseRepository) IsIPBanned(ip int) (bool, error) {
	database := r.DB
	var ipBlacklist models.IPBlacklist
	result := database.First(&ipBlacklist, ip)
	if result.Error == nil {
		return false, result.Error
	} else if result.Error == gorm.ErrRecordNotFound {
		return true, nil
	}
	return true, nil
}

// IsUserAgentBanned checks if a user agent is banned
func (r *DatabaseRepository) IsUserAgentBanned(userAgent string) (bool, error) {
	if userAgent == "" {
		return false, nil
	}
	database := r.DB
	var uaBlacklist models.UABlacklist
	result := database.Where("ua = ?", userAgent).First(&uaBlacklist)
	if result.Error == nil {
		return true, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, result.Error
}

// GetDailyTotal retrieves the total requests count for a specific day
func (r *DatabaseRepository) GetDailyTotal(datetime time.Time) (int64, error) {
	var dailyTotal int64
	database := r.DB
	dailyResult := database.Model(models.HourlyStat{}).
		Select("COALESCE(SUM(request_count), 0)").
		Where("time >= ? and time < ?", datetime, datetime.Add(24*time.Hour)).
		Scan(&dailyTotal)
	if err := dailyResult.Error; err != nil {
		return 0, err
	}
	return dailyTotal, nil
}

// GetCustomerTotal retrieves the total requests count and invalid requests count for a specific customer for a specific day
func (r *DatabaseRepository) GetCustomerTotal(datetime time.Time, customer int) (*models.CustomerStats, error) {
	customerStat := &models.CustomerStats{}
	database := r.DB
	customerResult := database.Model(models.HourlyStat{}).
		Select("COALESCE(SUM(request_count), 0) as total_count, COALESCE(SUM(invalid_count), 0) as total_invalid_count").
		Where("time >= ? and time < ? and customer_id = ?", datetime, datetime.Add(24*time.Hour), customer).
		Scan(customerStat)
	print(customerStat.TotalCount)
	if err := customerResult.Error; err != nil {
		return nil, err
	}
	return customerStat, nil
}

type CustomerStats struct {
	TotalCount        int
	TotalInvalidCount int
}
