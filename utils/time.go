package utils

import (
	"time"
)

func RoundDownToHour(t int64) time.Time {
	// round requests down to the hour. 14:55:55 would be truncated as 14:00:00
	unix_time := time.Unix(t, 0)
	return time.Date(unix_time.Year(), unix_time.Month(), unix_time.Day(), unix_time.Hour(), 0, 0, 0, unix_time.Location())
}

func ParseDateString(dateString string) (time.Time, error) {
	layout := "20060102"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
