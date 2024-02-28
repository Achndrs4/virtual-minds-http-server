package database

import (
	"errors"
	"log"
	"time"
	"virtualminds/http-server/models"

	"gorm.io/gorm"
)

type MockDatabaseInterface interface {
	InitializeConnection(std_logger *log.Logger) error
	UpdateStatsCount(request *models.CustomerRequest, isValid bool) error
	GetCustomer(customerID uint) (*models.Customer, error)
	GetCustomerTotal(datetime time.Time, customer int) (*models.CustomerStats, error)
	GetDailyTotal(datetime time.Time) (int64, error)
	IsIPValid(ip int) (bool, error)
	IsUserAgentValid(userAgent string) (bool, error)
	CloseAll() error
}

type MockDatabase struct {
	mockDB *gorm.DB
}

func (mdb *MockDatabase) InitializeConnection(std_logger *log.Logger) error {
	return nil
}

func (mdb *MockDatabase) UpdateStatsCount(request *models.CustomerRequest, isValid bool) error {
	if request.CustomerID == 1 {
		return nil
	}
	return errors.New("unhandled err")
}

func (mdb *MockDatabase) GetCustomer(customerID uint) (*models.Customer, error) {
	// Mock implementation to return a predefined customer
	if customerID == 1 {
		return &models.Customer{ID: customerID, Name: "", Active: true}, nil
	}
	if customerID == 2 {
		return &models.Customer{ID: customerID, Name: "", Active: false}, nil
	}
	return nil, errors.New("unhandled err")

}

func (mdb *MockDatabase) GetCustomerTotal(datetime time.Time, customer int) (*models.CustomerStats, error) {
	// Mock implementation to return a predefined customer stats
	if customer == 1 {
		return &models.CustomerStats{TotalCount: 10, TotalInvalidCount: 2}, nil
	}
	return nil, errors.New("unhandled err")
}

func (mdb *MockDatabase) GetDailyTotal(datetime time.Time) (int64, error) {
	// Mock implementation to return a predefined daily total
	epochTime := int64(10)
	if datetime.Before(time.Unix(epochTime+10, 0)) {
		return 0, errors.New("unhandled err")
	}
	return 100, nil

}

func (mdb *MockDatabase) IsIPValid(ip int) (bool, error) {
	// Mock implementation to always return false for IPBanned
	if ip == 100 {
		return false, errors.New("unhandled err")
	}
	return true, nil
}

func (mdb *MockDatabase) IsUserAgentValid(userAgent string) (bool, error) {
	// Mock implementation to always return false for UserAgentBanned
	// Mock implementation to always return false for IPBanned
	if userAgent == "google" {
		return false, errors.New("unhandled err")
	}
	return true, nil
}

func (mdb *MockDatabase) CloseAll() error {
	// Mock implementation to close database connections
	return nil
}
