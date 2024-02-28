package utils

import (
	"testing"
	"virtualminds/http-server/models"

	"github.com/stretchr/testify/assert"
)

func TestIsValidIPv4(t *testing.T) {
	// Test cases with ipv4 strings
	validIPs := []string{"192.168.0.1", "10.0.0.1", "172.16.0.1"}
	for _, ip := range validIPs {
		isValidIp := IsValidIP(ip)
		if !isValidIp {
			t.Errorf("IP Expected to be Valid")
			continue
		}
	}
}

func TestIsValidIPv6(t *testing.T) {
	// Test cases with valid ipv6 strings
	validIPs := []string{"2001:db8:3333:4444:5555:6666:7777:8888", "::1234:5678", "2001:db8::"}
	for _, ip := range validIPs {
		isValidIp := IsValidIP(ip)
		if !isValidIp {
			t.Errorf("IP Expected to be Valid")
			continue
		}
	}
}

func TestInValidIP(t *testing.T) {
	// Test cases with valid inputs
	validIPs := []string{"192.1680.1", "10.01", "bad.0.1"}
	for _, ip := range validIPs {
		isValidIp := IsValidIP(ip)
		if isValidIp {
			t.Errorf("IP Expected to be Valid")
			continue
		}
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
