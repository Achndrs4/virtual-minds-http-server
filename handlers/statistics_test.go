package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	"virtualminds/http-server/mocks"
	"virtualminds/http-server/services"
)

func setupStatsTestRouter() (*services.StatsService, *gin.Engine) {
	mockCustomerService := &services.StatsService{DB: &mocks.MockDatabase{}}
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return mockCustomerService, router
}

func TestGetDailyStats(t *testing.T) {
	mockStatsService, router := setupStatsTestRouter()
	// Register the GetDailyStats handler with the router
	router.GET("/statistics", func(c *gin.Context) {
		DailyStats(c, mockStatsService)
	})

	// Create a test request
	req, err := http.NewRequest("GET", "/statistics?customer=1&date=20240228", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponseBody := `{"Customer Invalid Requests":2,"Customer Valid Requests":10,"Total Daily Requests":100}`
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

func TestGetDailyStatsError(t *testing.T) {
	mockStatsService, router := setupStatsTestRouter()

	router.GET("/statistics", func(c *gin.Context) {
		DailyStats(c, mockStatsService)
	})

	// Create a test request with a bad request that the database will return an error (userNotFound)``
	req, err := http.NewRequest("GET", "/statistics?customer=2&date=20240228", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// that we properly handle errors
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
