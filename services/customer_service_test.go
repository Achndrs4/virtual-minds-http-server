package services

import (
	"testing"

	"virtualminds/http-server/mocks"
	"virtualminds/http-server/models"

	"github.com/stretchr/testify/assert"
)

var customer_service = CustomerService{DB: &mocks.MockDatabase{}}

func TestIsCustomerValid_ValidCustomer(t *testing.T) {

	// Call the IsCustomerValid function with a valid customer ID
	isValid, err := customer_service.IsCustomerValid(mocks.GOOD_CUSTOMER_ID)

	// Check the result
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsCustomerValid_InactiveCustomer(t *testing.T) {
	// Call the IsCustomerValid function with an inactive customer ID
	isValid, err := customer_service.IsCustomerValid(mocks.INACTIVE_CUSTOMER_ID)

	// Check the result
	assert.NoError(t, err)
	assert.False(t, isValid)
}

func TestIsCustomerValid_CustomerNotFound(t *testing.T) {
	// Call the IsCustomerValid function with a non-existing customer ID
	_, err := customer_service.IsCustomerValid(mocks.NONEXISTANT_CUSTOMER_ID)

	// Check the result
	assert.Error(t, err)
}

func TestValidUpdateCustomerStatistic(t *testing.T) {

	request := &models.CustomerRequest{
		CustomerID: mocks.GOOD_CUSTOMER_ID,
		RemoteIP:   mocks.GOOD_IP,
		Timestamp:  0,
	}

	err := customer_service.WriteCustomerStatistic(request, true)
	assert.NoError(t, err)
}

func TestInvalidUpdateCustomerStatistic(t *testing.T) {
	request := &models.CustomerRequest{
		CustomerID: mocks.NONEXISTANT_CUSTOMER_ID,
		RemoteIP:   mocks.GOOD_IP,
		Timestamp:  0,
	}

	err := customer_service.WriteCustomerStatistic(request, true)
	assert.Error(t, err)
}

func TestUnbannedIP(t *testing.T) {
	isIpBanned, err := customer_service.IsIPBanned(mocks.GOOD_IP)
	assert.NoError(t, err)
	assert.False(t, isIpBanned)
}

func TestBannedIP(t *testing.T) {
	isIpBanned, err := customer_service.IsIPBanned(mocks.BLOCKED_IP)
	assert.NoError(t, err)
	assert.True(t, isIpBanned)
}
