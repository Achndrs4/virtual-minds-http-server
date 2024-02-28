package utils

import (
	"strconv"
	"strings"
)

func GetValidIP(ip string) (int, error) {
	ip_int, err := strconv.Atoi(strings.ReplaceAll(ip, ".", ""))
	if err != nil {
		return -1, err
	}
	return ip_int, nil
}
