package services

import (
	"testing"
	"time"
	"virtualminds/http-server/database"

	"github.com/stretchr/testify/assert"
)

func TestValidGetCustomerTotal(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the WriteCustomerStatistic function with a good value
	time := time.Now()
	customer := 1
	count, err := service.DB.GetCustomerTotal(time, customer)
	assert.NoError(t, err)
	assert.Equal(t, count.TotalCount, int64(10))
	assert.Equal(t, count.TotalInvalidCount, int64(2))
}

func TestInvalidGetCustomerTotal(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the WriteCustomerStatistic function with a good value
	time := time.Now()
	customer := 1
	count, err := service.DB.GetCustomerTotal(time, customer)
	assert.NoError(t, err)
	assert.Equal(t, count.TotalCount, int64(10))
	assert.Equal(t, count.TotalInvalidCount, int64(2))
}

func TestInvalidGetDailyStats(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the GetDailyTotal function with a good value
	epoch := time.Unix(0, 0)
	_, err := service.DB.GetDailyTotal(epoch)
	assert.Error(t, err)
}

func TestValidGetDailyStatsl(t *testing.T) {
	service := CustomerService{DB: &database.MockDatabase{}}
	// Call the WriteCustomerStatistic function with a good value
	time := time.Now()
	total, err := service.DB.GetDailyTotal(time)
	assert.NoError(t, err)
	assert.Equal(t, total, int64(100))
}
