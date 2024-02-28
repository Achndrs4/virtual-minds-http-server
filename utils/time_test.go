package utils

import (
	"testing"
	"time"
)

func TestRoundDownToHour(t *testing.T) {
	// Test case with arbitrary time
	arbitraryTime := time.Date(2024, time.February, 28, 14, 55, 55, 0, time.UTC)
	expectedRoundedTime := time.Date(2024, time.February, 28, 14, 0, 0, 0, time.UTC)

	roundedTime := RoundDownToHour(arbitraryTime.Unix())

	if !roundedTime.Equal(expectedRoundedTime) {
		t.Errorf("RoundDownToHour failed for arbitrary time: expected %v, got %v", expectedRoundedTime, roundedTime)
	}

	// Test case with Unix epoch time (January 1, 1970)
	epochTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedRoundedTimeEpoch := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	roundedTimeEpoch := RoundDownToHour(epochTime.Unix())

	if !roundedTimeEpoch.Equal(expectedRoundedTimeEpoch) {
		t.Errorf("RoundDownToHour failed for epoch time: expected %v, got %v", expectedRoundedTimeEpoch, roundedTimeEpoch)
	}
}

func TestParseDateString(t *testing.T) {
	// Test case with valid date string
	validDateString := "20240228"
	expectedParsedTime := time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC)

	parsedTime, err := ParseDateString(validDateString)

	if err != nil {
		t.Errorf("ParseDateString failed for valid date string: %v", err)
	}

	if !parsedTime.Equal(expectedParsedTime) {
		t.Errorf("ParseDateString failed for valid date string: expected %v, got %v", expectedParsedTime, parsedTime)
	}

	// Test case with invalid date string
	invalidDateString := "2024-02-28" // invalid format
	_, err = ParseDateString(invalidDateString)

	if err == nil {
		t.Errorf("ParseDateString did not return error for invalid date string: %s", invalidDateString)
	}
}
