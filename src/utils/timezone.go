package utils

import (
	"fmt"
	"time"
)

func ConvertYearMonthDayToTimezone(yearMonthDay, timezone string) (time.Time, error) {
	// Load the timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load timezone: %v", err)
	}

	// Define a default day and time for cases where only year and month are given
	yearMonthDay = yearMonthDay + " 00:00:00" // Default time is midnight

	// Parse the input dateTime in the specified timezone
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", yearMonthDay, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date/time: %v", err)
	}

	return parsedTime, nil
}

// ConvertYearMonthDayToNextDay converts a date (year, month, and day) in a specific timezone to the next day's time.Time object
func ConvertYearMonthDayToNextDay(yearMonthDay, timezone string) (time.Time, error) {
	// Load the timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load timezone: %v", err)
	}

	// Add default time to the year, month, and day
	yearMonthDay = yearMonthDay + " 00:00:00" // Default time is midnight

	// Parse the input date in the specified timezone
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", yearMonthDay, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse year, month, and day: %v", err)
	}

	// Add one day to the parsed time
	nextDay := parsedTime.AddDate(0, 0, 1)

	return nextDay, nil
}

// ConvertToNextDay converts a date (year, month, and day) in a specific timezone to the next day's "year-month-day" string
func ConvertToNextDay(yearMonthDay, timezone string) (string, error) {
	// Load the timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", fmt.Errorf("failed to load timezone: %v", err)
	}

	// Add default time to the year, month, and day
	yearMonthDay = yearMonthDay + " 00:00:00" // Default time is midnight

	// Parse the input date in the specified timezone
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", yearMonthDay, loc)
	if err != nil {
		return "", fmt.Errorf("failed to parse year, month, and day: %v", err)
	}

	// Add one day to the parsed time
	nextDay := parsedTime.AddDate(0, 0, 1)

	// Format the next day as "year-month-day"
	nextDayFormatted := nextDay.Format("2006-01-02")

	return nextDayFormatted, nil
}

// ConvertBetweenTimezones converts a datetime string from one timezone to another
func ConvertBetweenTimezones(dateTime, fromTimezone, toTimezone string) (time.Time, error) {
	// Load the source timezone
	fromLoc, err := time.LoadLocation(fromTimezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load source timezone: %v", err)
	}

	// Load the target timezone
	toLoc, err := time.LoadLocation(toTimezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load target timezone: %v", err)
	}

	// Parse the input datetime in the source timezone
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateTime, fromLoc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime: %v", err)
	}

	// Convert the time to the target timezone
	convertedTime := parsedTime.In(toLoc)
	return convertedTime, nil
}
