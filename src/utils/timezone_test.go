package utils

import (
	"testing"
)

// go test -v -run TestConvertYearMonthDayToTimezone
func TestConvertYearMonthDayToTimezone(t *testing.T) {

	yearMonthDay := "2024-12-01" // Input year and month
	timezone := "UTC"            // Timezone (e.g., "UTC", "Europe/London")

	convertedTime, err := ConvertYearMonthDayToTimezone(yearMonthDay, timezone)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}

	t.Logf("Converted time in %s: %v\n", timezone, convertedTime)
}

// go test -v -run TestConvertYearMonthDayToNextDay
func TestConvertYearMonthDayToNextDay(t *testing.T) {

	yearMonthDay := "2024-12-01" // Input year and month
	timezone := "UTC"            // Timezone (e.g., "UTC", "Europe/London")

	nextDay, err := ConvertYearMonthDayToNextDay(yearMonthDay, timezone)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}

	t.Logf("Next day in %s: %v\n", timezone, nextDay)
}

// go test -v -run TestConvertBetweenTimezones
func TestConvertBetweenTimezones(t *testing.T) {

	dateTime := "2024-12-31 14:00:00" // Input datetime in source timezone
	fromTimezone := "Asia/Seoul"      // Source timezone
	toTimezone := "America/New_York"  // Target timezone

	convertedTime, err := ConvertBetweenTimezones(dateTime, fromTimezone, toTimezone)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}

	t.Logf("Converted time from %s to %s: %v\n", fromTimezone, toTimezone, convertedTime)
}
