package tracker

import (
	"context"
	"sync"
	"testing"
	"time"

	"token-tracker/configs"
	"token-tracker/get"
	"token-tracker/logger"
	"token-tracker/utils"

	"go.uber.org/zap"
)

// go test -v -timeout 30m -run TestEnhancedBlockTimestampRecorder
func TestEnhancedBlockTimestampRecorder(t *testing.T) {

	configs.SetEnv()

	filePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp.json"
	b, err := readJSONFile(filePath)
	if err != nil {
		t.Error(err)
	}

	timestampMap := mapTimestampByNumber(b)
	hexMap := mapHexTimestampByNumber(b)
	structMap := mapStructByNumber(b)
	t.Logf("Timestamps:\n %+v", timestampMap)

	// Separate variables for full file saving and partial file saving to handle different scenarios.
	var BlockTimestampMap sync.Map
	var BlockTimestampMap2 sync.Map
	for k, v := range hexMap {
		BlockTimestampMap.Store(k, v)
		BlockTimestampMap2.Store(k, v)
	}

	// currentBlockNumber := "0x14"
	currentBlockNumber, err := get.GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	// number of goroutine
	nog := 1
	numRecords := 10
	// channel
	isBlockWithDataChans := make([]chan map[string]struct{}, nog)
	blockTimestampMapChan := make(chan Task)
	errChan := make(chan error, nog)

	for i := 0; i < nog; i++ {
		isBlockWithDataChans[i] = make(chan map[string]struct{}, numRecords)

		// Launch a goroutine to collect timestamps.
		go EnhancedBlockTimestampRecorder(i, string(currentBlockNumber), numRecords,
			isBlockWithDataChans[i],
			blockTimestampMapChan, errChan)

		// Deliver initial data.
		go func() {
			isBlockWithDataChans[i] <- structMap
		}()
	}

	totalTime := 180
	intervalTime := 120
	// Declare ctx (context) and ticker.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(totalTime)*time.Minute)
	defer cancel()
	ticker := time.NewTicker(time.Duration(intervalTime) * time.Second)
	defer ticker.Stop()

	for {
		select {

		// Store collected timestamps in a map after reaching a certain amount
		// and pass them to another goroutine.
		case task := <-blockTimestampMapChan:

			records := make(map[string]struct{}, numRecords)
			for k, v := range task.HexTimestamps {
				BlockTimestampMap.Store(k, v)
				BlockTimestampMap2.Store(k, v)
				records[k] = struct{}{}
			}
			for i := 0; i < nog; i++ {
				if i != task.Id {
					isBlockWithDataChans[i] <- records
				}
			}

		case e := <-errChan:
			_ = e
			logger.Log.Warn("Too many requests to Infura", zap.Int("status_code", 429))

		case <-ticker.C:
			logger.Log.Infof("Ticker ticked: %d seconds passed\n", intervalTime)

			saved := make([]BlockTimestamp, 0)
			BlockTimestampMap2.Range(func(key, value interface{}) bool {
				logger.Log.Infof("Block: %v, Timestamp: %v\n", key, value)

				keyStr, _ := key.(string)
				valueStr, _ := value.(string)
				timestamp, _ := ConvertBlockTimestamp(valueStr)
				saved = append(saved, BlockTimestamp{Number: keyStr, Timestamp: timestamp})

				return true
			})

			timeFilePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp-temp.json"
			err = utils.EnsureFileExists(timeFilePath)
			if err != nil {
				t.Error("Error checking file: ", err)
			}
			// TODO: Compare unbuffered and buffered storage methods when data accumulates.
			err = utils.SaveJSONToFile(saved, timeFilePath)
			if err != nil {
				t.Error(err)
			}

		case <-ctx.Done():
			logger.Log.Info("Context timeout:", ctx.Err())

			saved := make([]BlockTimestamp, 0)
			BlockTimestampMap.Range(func(key, value interface{}) bool {
				logger.Log.Infof("Block: %v, Timestamp: %v\n", key, value)

				keyStr, _ := key.(string)
				valueStr, _ := value.(string)
				timestamp, _ := ConvertBlockTimestamp(valueStr)
				saved = append(saved, BlockTimestamp{Number: keyStr, Timestamp: timestamp})

				return true
			})

			timeFilePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp.json"
			err = utils.EnsureFileExists(timeFilePath)
			if err != nil {
				t.Error("Error checking file: ", err)
			}
			// TODO: Compare unbuffered and buffered storage methods when data accumulates.
			err = utils.SaveJSONToFile(saved, timeFilePath)
			if err != nil {
				t.Error(err)
			}

			t.Log("EnhancedBlockTimestampRecorder execution completed.")
			return
		}
	}
	// time.Sleep(time.Duration(math.MaxInt64))	// time.Sleep(100 * time.Second)
}
