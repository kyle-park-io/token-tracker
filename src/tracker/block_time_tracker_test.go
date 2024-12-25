package tracker

import (
	"testing"

	"token-tracker/configs"
	"token-tracker/get"
	"token-tracker/logger"
	"token-tracker/utils"

	"go.uber.org/zap"
)

// go test -v -run TestRecordBlockTimestamp
func TestRecordBlockTimestamp(t *testing.T) {

	configs.SetEnv()

	currentBlockNumber, err := get.GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	// number of goroutine
	nog := 3
	numBlocks := 10
	doneChan := make(chan struct{}, nog)
	errChan := make(chan error, nog)

	for i := 0; i < nog; i++ {
		go RecordBlockTimestamp(string(currentBlockNumber), int64(numBlocks),
			doneChan, errChan)
	}

	count := 0
	saved := make([]BlockTimestamp, 0, numBlocks)
forLoop:
	for {
		select {
		case <-doneChan:

			count++
			if count == nog {
				close(doneChan)
				close(errChan)

				// Retrieve and  all recorded block timestamps
				logger.Log.Info("Stored Block Timestamps: ")

				BlockTimestampMap.Range(func(key, value interface{}) bool {
					logger.Log.Infof("Block: %v, Timestamp: %v\n", key, value)

					keyStr, _ := key.(string)
					valueStr, _ := value.(string)
					timestamp, _ := ConvertBlockTimestamp(valueStr)
					saved = append(saved, BlockTimestamp{Number: keyStr, Timestamp: timestamp})

					return true
				})

				break forLoop
			}

		case e := <-errChan:
			_ = e
			logger.Log.Warn("Too many requests to Infura", zap.Int("status_code", 429))
		}
	}

	timeFilePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp-example.json"
	err = utils.EnsureFileExists(timeFilePath)
	if err != nil {
		t.Error("Error checking file: ", err)
	}
	err = utils.SaveJSONToFile(saved, timeFilePath)
	if err != nil {
		t.Error(err)
	}
}

// go test -v -run TestConvertBlockTimestamp
func TestConvertBlockTimestamp(t *testing.T) {

	configs.SetEnv()

	blockNumber, err := get.GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	hexTimestamp, err := get.GetBlockTimestampByNumber((string(blockNumber)))
	if err != nil {
		t.Error(err)
	}

	timestamp, err := ConvertBlockTimestamp(hexTimestamp)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Timestamp: %+v", timestamp)
}