package utils

import (
	"encoding/json"
	"net"
	"virtualminds/http-server/models"
)

func IsValidIP(ip string) bool {
	// check if the IP is valid. Accepts IPv4 dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"), or IPv4-mapped IPv6
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

func GetRequestBody(body []byte) (*models.CustomerRequest, error) {
	// check if the json can be parsed, and if it has the required fields
	var request models.CustomerRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return &request, nil
}
