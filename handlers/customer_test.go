package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"virtualminds/http-server/mocks"
	"virtualminds/http-server/models"
	"virtualminds/http-server/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupCustomerTestRouter() (*services.CustomerService, *gin.Engine) {
	// set up a router and a mock service
	mockCustomerService := &services.CustomerService{DB: &mocks.MockDatabase{}}
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return mockCustomerService, router
}

func TestPersistCustomerEntryGoodData(t *testing.T) {
	mockCustomerService, router := setupCustomerTestRouter()

	good_data := &models.CustomerRequest{
		CustomerID: mocks.GOOD_CUSTOMER_ID,
		RemoteIP:   mocks.GOOD_IP,
		Timestamp:  0,
	}
	jsonData, _ := json.Marshal(good_data)

	// Register the GetDailyStats handler with the router
	router.POST("/customer", func(c *gin.Context) {
		c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
		PersistCustomerEntry(c, mockCustomerService)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPersistCustomerIDNotFound(t *testing.T) {
	mockCustomerService, router := setupCustomerTestRouter()

	good_data := &models.CustomerRequest{
		CustomerID: 2,
		RemoteIP:   "53.135.119.210",
		Timestamp:  100,
	}
	jsonData, _ := json.Marshal(good_data)

	// Register the GetDailyStats handler with the router
	router.POST("/customer", func(c *gin.Context) {
		c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
		PersistCustomerEntry(c, mockCustomerService)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPersistCustomerInactive(t *testing.T) {
	mockCustomerService, router := setupCustomerTestRouter()
	good_data := &models.CustomerRequest{
		CustomerID: 3,
		RemoteIP:   "53.135.119.210",
		Timestamp:  100,
	}
	jsonData, _ := json.Marshal(good_data)

	// Register the GetDailyStats handler with the router
	router.POST("/customer", func(c *gin.Context) {
		c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
		PersistCustomerEntry(c, mockCustomerService)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPersistCustomerEntryBlockedIP(t *testing.T) {
	mockCustomerService, router := setupCustomerTestRouter()

	good_data := &models.CustomerRequest{
		CustomerID: mocks.GOOD_CUSTOMER_ID,
		RemoteIP:   mocks.BLOCKED_IP,
		Timestamp:  0,
	}
	jsonData, _ := json.Marshal(good_data)

	// Register the GetDailyStats handler with the router
	router.POST("/customer", func(c *gin.Context) {
		c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
		PersistCustomerEntry(c, mockCustomerService)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestPersistCustomerEntryBlockedAgent(t *testing.T) {
	mockCustomerService, router := setupCustomerTestRouter()

	good_data := &models.CustomerRequest{
		CustomerID: mocks.GOOD_CUSTOMER_ID,
		RemoteIP:   mocks.GOOD_IP,
		Timestamp:  0,
	}
	jsonData, _ := json.Marshal(good_data)

	// Register the GetDailyStats handler with the router
	router.POST("/customer", func(c *gin.Context) {
		c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
		c.Request.Header.Set("User-Agent", mocks.BAD_USER_AGENT)
		PersistCustomerEntry(c, mockCustomerService)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/customer", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a test HTTP recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusForbidden, w.Code)
}
