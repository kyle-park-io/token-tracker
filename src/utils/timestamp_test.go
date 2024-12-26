package utils

import (
	"testing"
)

// go test -v -run TestHexToKST
func TestHexToKST(t *testing.T) {
	// Example usage of the functions
	hex := "0x67694603"

	// Convert hex to KST time
	kstTime, err := HexToKST(hex)
	if err != nil {
		t.Error("Error converting hex to KST:", err)
	}
	t.Log("KST Time:", kstTime)
}

// go test -v -run TestHexToUnix
func TestHexToUnix(t *testing.T) {
	// Example usage of the functions
	hex := "0x67694603"

	// Convert hex to Unix timestamp
	unixTimestamp, err := HexToUnix(hex)
	if err != nil {
		t.Error("Error converting hex to Unix:", err)
	}
	t.Log("Unix Timestamp:", unixTimestamp)
}

// go test -v -run TestUnixToTime
func TestUnixToTime(t *testing.T) {
	// Example usage of the functions
	hex := "0x67694603"
	timezone := "America/New_York"

	// Convert hex to Unix timestamp
	unixTimestamp, err := HexToUnix(hex)
	if err != nil {
		t.Error("Error converting hex to Unix:", err)
	}

	// Convert Unix timestamp to time in a specific timezone
	localTime, err := UnixToTime(unixTimestamp, timezone)
	if err != nil {
		t.Error("Error converting Unix to time in timezone:", err)
	}
	t.Log("Time in New York:", localTime)
}

// go test -v -run TestTimeToUnix
func TestTimeToUnix(t *testing.T) {
	// Example usage of the functions
	yearMonthDay := "2024-12-26"
	timezone := "UTC"

	// Convert date to timezone
	convertedTime, err := ConvertYearMonthDayToTimezone(yearMonthDay, timezone)
	if err != nil {
		t.Error("Error converting date to timezone:", err)
	}

	// Convert Unix timestamp to time in a specific timezone
	unixTime, err := TimeToUnix(convertedTime, timezone)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	t.Logf("Unix timestamp for %v in %s: %d\n", convertedTime, timezone, unixTime)
}
