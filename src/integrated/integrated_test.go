package integrated

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/tracker"
	"github.com/kyle-park-io/token-tracker/utils"
)

// go test -v -run TestFindBlockByTimestamp
func TestFindBlockByTimestamp(t *testing.T) {

	// Initialize configuration environment
	config.SetDevEnv()

	// Test input values
	yearMonthDay := "2024-12-24"
	timezone := "UTC"

	// 1. Convert the given date to the next day
	nextDay, err := utils.ConvertYearMonthDayToNextDay(yearMonthDay, timezone)
	if err != nil {
		t.Errorf("Failed to convert to next day: %v", err)
	}

	// 2. Convert the adjusted date to a Unix timestamp
	unixTime, err := utils.TimeToUnix(nextDay, timezone)
	if err != nil {
		t.Errorf("Failed to convert time to Unix: %v", err)
	}
	t.Logf("Unix timestamp for %v in %s: %d", nextDay, timezone, unixTime)

	// 3. Track the block range based on the Unix timestamp
	blockRange, err := tracker.TrackBlockTimestamp(unixTime)
	if err != nil {
		t.Errorf("Failed to track block timestamp: %v", err)
	}
	t.Logf("Block range: Low=%d, High=%d", blockRange.Low, blockRange.High)

	// 4. Convert block range values to hexadecimal format
	lowBlockHex := utils.DecimalToHex(blockRange.Low)
	highBlockHex := utils.DecimalToHex(blockRange.High)

	// 5. Retrieve block timestamps using block numbers
	lowBlockTimestampHex, err := get.GetBlockTimestampByNumber(lowBlockHex)
	if err != nil {
		t.Errorf("Failed to get block timestamp for low block: %v", err)
	}

	highBlockTimestampHex, err := get.GetBlockTimestampByNumber(highBlockHex)
	if err != nil {
		t.Errorf("Failed to get block timestamp for high block: %v", err)
	}

	// 6. Convert retrieved timestamps to Unix format
	lowBlockUnix, err := utils.HexToUnix(lowBlockTimestampHex)
	if err != nil {
		t.Errorf("Failed to convert low block timestamp to Unix: %v", err)
	}

	highBlockUnix, err := utils.HexToUnix(highBlockTimestampHex)
	if err != nil {
		t.Errorf("Failed to convert high block timestamp to Unix: %v", err)
	}

	// 7. Output results
	t.Logf("Low Block Unix Timestamp: %d", lowBlockUnix)
	t.Logf("High Block Unix Timestamp: %d", highBlockUnix)
	t.Logf("Input Unix Timestamp: %d", unixTime)
}
