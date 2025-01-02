package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/tracker"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/spf13/viper"
)

func EnhancedBlockTimestampRecorder() {

	// time.Sleep(time.Duration(math.MaxInt64))

	fileName := "blockTimestamp.json"
	folderPath := viper.GetString("ROOT_PATH") + "/json/blockTimestamp"
	filePath := folderPath + "/" + fileName

	err := utils.CreateFolder(folderPath)
	if err != nil {
		logger.Log.Warnln(err)
	}
	err = utils.EnsureFileExists(filePath)
	if err != nil {
		logger.Log.Warnln(err)
	}

	b, err := tracker.ReadBlockTimestampJSONFile(filePath)
	if err != nil {
		logger.Log.Errorln(err)
	}

	timestampMap := tracker.MapTimestampByNumber(b)
	hexMap := tracker.MapHexTimestampByNumber(b)
	structMap := tracker.MapStructByNumber(b)
	logger.Log.Infof("Timestamps:\n %+v", timestampMap)

	// Separate variables for full file saving and partial file saving to handle different scenarios.
	var BlockTimestampMap sync.Map
	var BlockTimestampMap2 sync.Map
	for k, v := range hexMap {
		BlockTimestampMap.Store(k, v)
		BlockTimestampMap2.Store(k, v)
	}

	currentBlockNumber, err := get.GetBlockNumber()
	if err != nil {
		logger.Log.Errorln(err)
		logger.Log.Fatalln(err)
	}

	// number of goroutine
	nog := 1
	numRecords := 30
	// channel
	isBlockWithDataChans := make([]chan map[string]struct{}, nog)
	blockTimestampMapChan := make(chan tracker.Task)
	errChan := make(chan error, nog)

	for i := 0; i < nog; i++ {
		isBlockWithDataChans[i] = make(chan map[string]struct{}, numRecords)

		// Launch a goroutine to collect timestamps.
		go tracker.EnhancedBlockTimestampRecorder(i, string(currentBlockNumber), numRecords,
			isBlockWithDataChans[i],
			blockTimestampMapChan, errChan)

		// Deliver initial data.
		go func() {
			isBlockWithDataChans[i] <- structMap
		}()
	}

	totalTime := 3 * 60 * 60
	// totalTime := math.MaxInt64
	intervalTime := 300
	// Declare ctx (context) and ticker.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(totalTime)*time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(totalTime)*time.Second)
	defer cancel()
	ticker := time.NewTicker(time.Duration(intervalTime) * time.Second)
	defer ticker.Stop()

	count := 0
	for {
		select {

		// Store collected timestamps in a map after reaching a certain amount
		// and pass them to another goroutine.
		case task := <-blockTimestampMapChan:

			records := make(map[string]struct{}, numRecords)
			for k, v := range task.HexTimestamps {
				count++
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
			// logger.Log.Warnln("Too many requests to Infura", zap.Int("status_code", 429))
			logger.Log.Warnln(e.Error())
			// ws.GlobalLogChannel <- e.Error()
			wss.GlobalLogChannel <- e.Error()

		case <-ticker.C:
			logger.Log.Infof("Ticker ticked: %d seconds passed\n", intervalTime)

			saved := make([]tracker.BlockTimestamp, 0)
			BlockTimestampMap2.Range(func(key, value interface{}) bool {
				log := fmt.Sprintf("Block: %v, Timestamp: %v\n", key, value)
				logger.Log.Infof(log)

				keyStr, _ := key.(string)
				valueStr, _ := value.(string)
				timestamp, _ := tracker.ConvertBlockTimestamp(valueStr)
				saved = append(saved, tracker.BlockTimestamp{Number: keyStr, Timestamp: timestamp})

				return true
			})

			timeFilePath := viper.GetString("ROOT_PATH") + "/json/blockTimestamp/blockTimestamp-temp.json"
			timeFilePath2 := viper.GetString("ROOT_PATH") + "/json/blockTimestamp/blockTimestamp.json"
			err = utils.EnsureFileExists(timeFilePath)
			if err != nil {
				logger.Log.Errorln("Error checking file: ", err)
			}
			// TODO: Compare unbuffered and buffered storage methods when data accumulates.
			err = utils.SaveJSONToFile(saved, timeFilePath)
			if err != nil {
				logger.Log.Errorln(err)
			}
			err = utils.SaveJSONToFile(saved, timeFilePath2)
			if err != nil {
				logger.Log.Errorln(err)
			}

			// pvc
			logger.Log.Infoln("Save to PVC")
			if viper.GetString("ENV") == "prod" {
				timeFilePath3 := viper.GetString("ROOT_PATH") + "/../data/blockTimestamp.json"
				err = utils.SaveJSONToFile(saved, timeFilePath3)
				if err != nil {
					logger.Log.Errorln(err)
				}
			}

		case <-ctx.Done():
			logger.Log.Info("Context timeout:", ctx.Err())

			saved := make([]tracker.BlockTimestamp, 0)
			BlockTimestampMap.Range(func(key, value interface{}) bool {
				logger.Log.Infof("Block: %v, Timestamp: %v\n", key, value)

				keyStr, _ := key.(string)
				valueStr, _ := value.(string)
				timestamp, _ := tracker.ConvertBlockTimestamp(valueStr)
				saved = append(saved, tracker.BlockTimestamp{Number: keyStr, Timestamp: timestamp})

				return true
			})

			timeFilePath := viper.GetString("ROOT_PATH") + "/json/blockTimestamp/blockTimestamp.json"
			err = utils.EnsureFileExists(timeFilePath)
			if err != nil {
				logger.Log.Errorln("Error checking file: ", err)
			}
			// TODO: Compare unbuffered and buffered storage methods when data accumulates.
			err = utils.SaveJSONToFile(saved, timeFilePath)
			if err != nil {
				logger.Log.Errorln(err)
			}

			// pvc
			logger.Log.Infoln("Save to PVC")
			if viper.GetString("ENV") == "prod" {
				timeFilePath2 := viper.GetString("ROOT_PATH") + "/../data/blockTimestamp.json"
				err = utils.SaveJSONToFile(saved, timeFilePath2)
				if err != nil {
					logger.Log.Errorln(err)
				}
			}

			logger.Log.Infoln("Total count: ", count)
			logger.Log.Infoln("EnhancedBlockTimestampRecorder execution completed.")
			return
		}
	}
}
