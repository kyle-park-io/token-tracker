package tracker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"
)

type BlockTimestamp struct {
	Number    string    `json:"number"`
	Timestamp Timestamp `json:"timestamp"`
}

type Timestamp struct {
	Hex       string    `json:"hex"`
	Unix      int64     `json:"unix"`
	LocalTime time.Time `json:"localTime"`
}

var count int64
var BlockTimestampMap sync.Map

func RecordBlockTimestamp(currentBlockNumber string, numBlocks int64,
	doneChan chan<- struct{}, errorChan chan<- error) {

	for {
		if count >= numBlocks {
			doneChan <- struct{}{}
			return
		}

		// Generate a random hexadecimal string less than the latest block number.
		r, err := utils.RandomHexBelow(currentBlockNumber)
		if err != nil {
			errorChan <- err
		}

		logger.Log.Infof("Current number of recorded block: %d\n", count)
		t, err := get.GetBlockTimestampByNumber(r)
		if err != nil {
			errorChan <- err
			time.Sleep(1 * time.Second)
			continue
		}

		BlockTimestampMap.Store(r, t)
		atomic.AddInt64(&count, 1)
	}
}

type Task struct {
	Id            int               `json:"id"`
	HexTimestamps map[string]string `json:"hexTimestamps"`
}

func EnhancedBlockTimestampRecorder(id int, currentBlockNumber string, numRecords int,
	isBlockWithDataChan <-chan map[string]struct{},
	blockTimestampMapChan chan<- Task, errorChan chan<- error) {

	var initFn bool
	var internalCount int64
	var recordCount int
	var IsBlockWithDataMap sync.Map
	blockTimestampMap := make(map[string]string, numRecords)

	for {
		select {

		// Executes the function with the highest priority,
		// providing a performance advantage during execution.
		case m := <-isBlockWithDataChan:
			// TODO: Consider whether the counting process needs to be absolutely accurate.
			for k, _ := range m {
				IsBlockWithDataMap.Store(k, struct{}{})
				atomic.AddInt64(&internalCount, 1)
			}
			initFn = true

		default:

			// Check whether the prioritized execution element has been executed.
			if !initFn {
				continue
			}
			// time.Sleep(time.Duration(math.MaxInt64))

			// Generate a random hexadecimal string less than the latest block number.
			r, err := utils.RandomHexBelow(currentBlockNumber)
			if err != nil {
				logger.Log.Warnf("random error?")
				errorChan <- err
			}

			// Use sync.Map to check for duplicate keys.
			_, ok := IsBlockWithDataMap.Load(r)
			if ok {
				logger.Log.Infof("Already have blockTimestamp, block: %s\n", r)
				// time.Sleep(3 * time.Second)
				continue
			}

			// TODO: Verify if introducing a time sleep upon receiving a 429 error from Infura can improve performance.
			t, err := get.GetBlockTimestampByNumber(r)
			if err != nil {
				errorChan <- err
				// time.Sleep(3 * time.Second)
				continue
			}

			log := fmt.Sprintf("ID: %d, Current number of recorded block: %d", id, internalCount)
			logger.Log.Infoln(log)
			// ws.GlobalLogChannel <- log
			wss.GlobalLogChannel <- log

			blockTimestampMap[r] = t
			recordCount++
			IsBlockWithDataMap.Store(r, struct{}{})
			atomic.AddInt64(&internalCount, 1)

			// Send data to the main channel when the required number of records to log in the file is met.
			if recordCount == numRecords {
				task := Task{Id: id, HexTimestamps: blockTimestampMap}
				blockTimestampMapChan <- task

				// Initialize variables.
				blockTimestampMap = make(map[string]string, numRecords)
				recordCount = 0
			}

			time.Sleep(10 * time.Second)
		}
	}
}

func ConvertBlockTimestamp(hex string) (Timestamp, error) {

	unix, err := utils.HexToUnix(hex)
	if err != nil {
		return Timestamp{}, err
	}
	localTime, err := utils.UnixToTime(unix, "Asia/Seoul")
	if err != nil {
		return Timestamp{}, err
	}

	t := Timestamp{hex, unix, localTime}
	return t, nil
}
