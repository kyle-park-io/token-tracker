package utils

import (
	"fmt"
	"strconv"
	"time"
)

// HexToKST converts a hexadecimal string to a time in the Korean Standard Time (KST) zone.
func HexToKST(hex string) (time.Time, error) {
	// 1. Convert the hexadecimal string to a decimal Unix timestamp
	timestamp, err := strconv.ParseInt(hex, 0, 64) // Example: "0x67694603" → 1736505603
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse hex: %v", err)
	}

	// 2. Convert the Unix timestamp to UTC time
	utcTime := time.Unix(timestamp, 0)

	// 3. Convert the UTC time to Korean Standard Time (KST)
	kstLocation, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load timezone: %v", err)
	}
	kstTime := utcTime.In(kstLocation)

	return kstTime, nil
}

// HexToUnix converts a hexadecimal string to a Unix timestamp (decimal).
func HexToUnix(hex string) (int64, error) {
	// Convert the hexadecimal string to an integer
	timestamp, err := strconv.ParseInt(hex, 0, 64) // Example: "0x67694603" → 1736505603
	if err != nil {
		return 0, fmt.Errorf("failed to parse hex: %v", err)
	}
	return timestamp, nil
}

// UnixToTime converts a Unix timestamp to a time object in the specified timezone.
func UnixToTime(timestamp int64, timezone string) (time.Time, error) {
	// Load the specified timezone location
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load timezone: %v", err)
	}

	// Convert the Unix timestamp to time in the given timezone
	return time.Unix(timestamp, 0).In(loc), nil
}

// TimeToUnix converts a time object to a Unix timestamp in the specified timezone.
func TimeToUnix(t time.Time, timezone string) (int64, error) {
	// Load the specified timezone location
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, fmt.Errorf("failed to load timezone: %v", err)
	}

	// Convert the time to the specified timezone
	localTime := t.In(loc)

	// Return the Unix timestamp
	return localTime.Unix(), nil
}
