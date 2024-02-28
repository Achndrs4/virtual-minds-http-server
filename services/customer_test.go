package services

import (
	"testing"

	"virtualminds/http-server/database"
	"virtualminds/http-server/models"

	"github.com/stretchr/testify/assert"
)

func TestIsCustomerValid_ValidCustomer(t *testing.T) {
	// Create a new instance of the CustomerService with the fake database
	service := CustomerService{DB: &database.MockDatabase{}}

	// Call the IsCustomerValid function with a valid customer ID
	isValid, err := service.IsCustomerValid(1)

	// Check the result
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsCustomerValid_InactiveCustomer(t *testing.T) {
	// Create a new instance of the CustomerService with the fake database
	service := CustomerService{DB: &database.MockDatabase{}}

	// Call the IsCustomerValid function with an inactive customer ID
	isValid, err := service.IsCustomerValid(2)

	// Check the result
	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestIsCustomerValid_CustomerNotFound(t *testing.T) {
	// Create a new instance of the CustomerService with the fake database
	service := CustomerService{DB: &database.MockDatabase{}}

	// Call the IsCustomerValid function with a non-existing customer ID
	_, err := service.IsCustomerValid(3)

	// Check the result
	assert.Error(t, err)
}

func TestValidUpdateCustomerStatistic(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the WriteCustomerStatistic function with a good value
	request := &models.CustomerRequest{
		CustomerID: 1,
		RemoteIP:   "100",
		Timestamp:  100,
	}

	err := service.WriteCustomerStatistic(request, true)
	assert.NoError(t, err)
}

func TestInValidUpdateCustomerStatistic(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the WriteCustomerStatistic function with a good value
	request := &models.CustomerRequest{
		CustomerID: 2,
		RemoteIP:   "100",
		Timestamp:  100,
	}

	err := service.WriteCustomerStatistic(request, true)
	assert.Error(t, err)
}

func TestValidIP(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	isIpValid, err := service.DB.IsIPValid(99)
	assert.Equal(t, isIpValid, true)
	assert.NoError(t, err)
}

func TestInvalidIP(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	isIpValid, err := service.DB.IsIPValid(100)
	assert.Equal(t, isIpValid, false)
	assert.Error(t, err)
}
