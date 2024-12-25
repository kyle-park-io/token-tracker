package utils

import (
	"fmt"
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
		fmt.Println("Error converting Unix to time in timezone:", err)
		return
	}
	t.Log("Time in New York:", localTime)
}
