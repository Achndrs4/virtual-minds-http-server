package mocks

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
	IsUserAgentBanned(userAgent string) (bool, error)
	IsIPBanned(ip string) (bool, error)
	CloseAll() error
}

type MockDatabase struct {
	mockDB *gorm.DB
}

func (mdb *MockDatabase) InitializeConnection(std_logger *log.Logger) error {
	return nil
}

func (mdb *MockDatabase) UpdateStatsCount(request *models.CustomerRequest, isValid bool) error {
	if request.CustomerID == GOOD_CUSTOMER_ID {
		return nil
	}
	return errors.New("unhandled err")
}

func (mdb *MockDatabase) GetCustomer(customerID uint) (*models.Customer, error) {
	// Mock implementation to return a predefined customer
	if customerID == GOOD_CUSTOMER_ID {
		return &models.Customer{ID: customerID, Name: "", Active: true}, nil
	}
	if customerID == INACTIVE_CUSTOMER_ID {
		return &models.Customer{ID: customerID, Name: "", Active: false}, nil
	}
	return nil, errors.New("unhandled err")

}

func (mdb *MockDatabase) GetCustomerTotal(datetime time.Time, customer int) (*models.CustomerStats, error) {
	// Mock implementation to return a predefined customer stats
	if customer == GOOD_CUSTOMER_ID {
		return &models.CustomerStats{TotalCount: 10, TotalInvalidCount: 2}, nil
	}
	return nil, errors.New("unhandled err")
}

func (mdb *MockDatabase) GetDailyTotal(datetime time.Time) (int64, error) {
	// Nothing to mock here because this is a pure database function
	return 100, nil
}

func (mdb *MockDatabase) IsUserAgentBanned(userAgent string) (bool, error) {
	// Mock implementation to always return false for UserAgentBanned
	// Mock implementation to always return false for IPBanned
	if userAgent == BAD_USER_AGENT {
		return true, nil
	}
	return false, nil
}

func (mdb *MockDatabase) IsIPBanned(ip string) (bool, error) {
	// Mock implementation to always return false for UserAgentBanned
	// Mock implementation to always return false for IPBanned
	if ip == BLOCKED_IP {
		return true, nil
	}
	return false, nil
}

func (mdb *MockDatabase) CloseAll() error {
	// Mock implementation to close database connections
	return nil
}
