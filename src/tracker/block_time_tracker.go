package tracker

import (
	"math"

	"token-tracker/get"
	"token-tracker/logger"
	"token-tracker/utils"
)

type BlockPosition struct {
	Low  int64 `json:"low"`
	High int64 `json:"high"`
}

// Use binary search to find the position of the block.
func TrackBlockTimestamp(targetTime int64) (BlockPosition, error) {

	startBlock := "0x0"
	endBlock, err := get.GetBlockNumber()
	if err != nil {
		return BlockPosition{}, err
	}

	low, err := utils.HexToDecimal(startBlock)
	if err != nil {
		return BlockPosition{}, err
	}
	high, err := utils.HexToDecimal(string(endBlock))
	if err != nil {
		return BlockPosition{}, err
	}

	// Binary Search
	for {
		mid := (low + high) / 2
		midBlock := utils.DecimalToHex(mid)

		t, err := get.GetBlockTimestampByNumber(midBlock)
		if err != nil {
			return BlockPosition{}, err

		}
		timestamp, err := utils.HexToUnix(t)
		if err != nil {
			return BlockPosition{}, err
		}

		if timestamp == targetTime || math.Abs(float64(high)-float64(low)) == float64(1) {
			logger.Log.Info("The timestamp has been found.")

			if low <= high {
				b := BlockPosition{Low: low, High: high}
				return b, nil
			} else {
				b := BlockPosition{Low: high, High: low}
				return b, nil
			}
		} else if timestamp < targetTime {
			logger.Log.Info("The median value is smaller than the target time.")
			logger.Log.Infof("low: %d, high: %d\n", low, high)

			low = mid + 1
		} else {
			logger.Log.Info("The median value is bigger than the target time.")
			logger.Log.Infof("low: %d, high: %d\n", low, high)
			high = mid - 1
		}
	}
}
