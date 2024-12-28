package integrated

import (
	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/tracker"
	"github.com/kyle-park-io/token-tracker/utils"
)

func GetBlockPosition(yearMonthDay string, timezone string) (tracker.BlockPosition, error) {
	// 1. Convert the given date to the day
	day, err := utils.ConvertYearMonthDayToTimezone(yearMonthDay, timezone)
	if err != nil {
		logger.Log.Errorf("Failed to convert to next day: %v", err)
	}

	// 2. Convert the adjusted date to a Unix timestamp
	unixTime, err := utils.TimeToUnix(day, timezone)
	if err != nil {
		logger.Log.Errorf("Failed to convert time to Unix: %v", err)
	}
	logger.Log.Infof("Unix timestamp for %v in %s: %d", day, timezone, unixTime)

	// 3. Track the block range based on the Unix timestamp
	blockRange, err := tracker.TrackBlockTimestamp(unixTime)
	if err != nil {
		logger.Log.Errorf("Failed to track block timestamp: %v", err)
	}
	logger.Log.Infof("Block range: Low=%d, High=%d", blockRange.Low, blockRange.High)

	// 4. Convert block range values to hexadecimal format
	lowBlockHex := utils.DecimalToHex(blockRange.Low)
	highBlockHex := utils.DecimalToHex(blockRange.High)

	// 5. Retrieve block timestamps using block numbers
	lowBlockTimestampHex, err := get.GetBlockTimestampByNumber(lowBlockHex)
	if err != nil {
		logger.Log.Errorf("Failed to get block timestamp for low block: %v", err)
	}

	highBlockTimestampHex, err := get.GetBlockTimestampByNumber(highBlockHex)
	if err != nil {
		logger.Log.Errorf("Failed to get block timestamp for high block: %v", err)
	}

	// 6. Convert retrieved timestamps to Unix format
	lowBlockUnix, err := utils.HexToUnix(lowBlockTimestampHex)
	if err != nil {
		logger.Log.Errorf("Failed to convert low block timestamp to Unix: %v", err)
	}

	highBlockUnix, err := utils.HexToUnix(highBlockTimestampHex)
	if err != nil {
		logger.Log.Errorf("Failed to convert high block timestamp to Unix: %v", err)
	}

	// 7. Output results
	logger.Log.Infof("Low Block Unix Timestamp: %d", lowBlockUnix)
	logger.Log.Infof("High Block Unix Timestamp: %d", highBlockUnix)
	logger.Log.Infof("Input Unix Timestamp: %d", unixTime)

	return tracker.BlockPosition{Low: blockRange.Low, LowHex: lowBlockHex, LowTimestamp: lowBlockUnix, LowTimestampHex: lowBlockTimestampHex,
		High: blockRange.High, HighHex: highBlockHex, HighTimestamp: highBlockUnix, HighTimestampHex: highBlockTimestampHex}, nil
}
