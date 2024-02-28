package utils

import (
	"encoding/json"
	"strconv"
	"strings"
	"virtualminds/http-server/models"
)

func GetValidIP(ip string) (int, error) {
	ip_int, err := strconv.Atoi(strings.ReplaceAll(ip, ".", ""))
	if err != nil {
		return -1, err
	}
	return ip_int, nil
}

func GetRequestBody(body []byte) (*models.CustomerRequest, error) {
	// check if the json can be parsed, and if it has the required fields
	var request models.CustomerRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return &request, nil
}
