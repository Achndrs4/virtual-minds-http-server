package services

import (
	"testing"
	"time"
	"virtualminds/http-server/mocks"
	"virtualminds/http-server/utils"

	"github.com/stretchr/testify/assert"
)

var statistics_service = StatsService{DB: &mocks.MockDatabase{}}

func TestStatsServiceValidValidateRequest(t *testing.T) {
	// test if a valid request would work
	customer := mocks.VALID_CUSTOMER_STR
	dateString := mocks.VALID_DATE_STR
	expectedCustomerID := mocks.GOOD_CUSTOMER_ID
	expectedDate, _ := utils.ParseDateString(dateString)

	customerID, date, err := statistics_service.ValidateRequest(customer, dateString)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if customerID != expectedCustomerID {
		t.Errorf("Expected customer ID %d, got %d", expectedCustomerID, customerID)
	}
	if !date.Equal(*expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, date)
	}
}

func TestStatsServiceValidateRequestWithInvalidCustomer(t *testing.T) {
	// test if an invalid customer would work
	customer := mocks.INVALID_CUSTOMER_STR
	dateString := mocks.VALID_DATE_STR
	_, _, err := statistics_service.ValidateRequest(customer, dateString)
	if err == nil {
		t.Error("Expected error for invalid customer ID, got nil")
	}
}

func TestStatsServiceValidateRequestWithInvalidDate(t *testing.T) {
	// test if an invalid date format would work
	customer := mocks.VALID_CUSTOMER_STR
	dateString := mocks.INVALID_DATE_STR
	_, _, err := statistics_service.ValidateRequest(customer, dateString)
	if err == nil {
		t.Error("Expected error for invalid customer ID, got nil")
	}
}

func TestValidGetCustomerTotal(t *testing.T) {
	// Call the WriteCustomerStatistic function with a good value
	time := time.Now()
	customer := mocks.GOOD_CUSTOMER_ID
	count, err := statistics_service.DB.GetCustomerTotal(time, customer)
	assert.NoError(t, err)
	assert.Equal(t, count.TotalCount, int64(10))
	assert.Equal(t, count.TotalInvalidCount, int64(2))
}

func TestValidGetDailyStatsl(t *testing.T) {
	// Call the WriteCustomerStatistic function with a good value
	time := time.Now()
	total, err := statistics_service.DB.GetDailyTotal(time)
	assert.NoError(t, err)
	assert.Equal(t, total, int64(100))
}
