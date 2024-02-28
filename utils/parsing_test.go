package utils

import (
	"testing"
	"virtualminds/http-server/models"

	"github.com/stretchr/testify/assert"
)

func TestGetValidIP(t *testing.T) {
	// Test cases with valid inputs
	validIPs := []string{"192.168.0.1", "10.0.0.1", "172.16.0.1"}
	for _, ip := range validIPs {
		_, err := GetValidIP(ip)
		if err != nil {
			t.Errorf("Unexpected error for valid IP %s: %v", ip, err)
			continue
		}
	}

	// Test case with invalid input
	invalidIP := "invalid_ip"
	_, err := GetValidIP(invalidIP)
	if err == nil {
		t.Error("Expected error for invalid IP, but got nil")
	}
}

func TestGetRequestBody(t *testing.T) {
	// Test case: Valid JSON body
	validBody := []byte(`{"RemoteIP":"100","CustomerID":100,"Timestamp":1614620400}`)
	expectedRequest := &models.CustomerRequest{
		CustomerID: 100,
		RemoteIP:   "100",
		Timestamp:  1614620400,
	}

	request, err := GetRequestBody(validBody)
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedRequest, request, "Expected equal CustomerRequest")

	// Test case: Invalid JSON body
	invalidBody := []byte(`{"ID":1,"CustomerID":100,"Timestamp":"invalid"}`)
	_, err = GetRequestBody(invalidBody)
	assert.Error(t, err, "Expected an error")
}
