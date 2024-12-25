package tracker

import (
	"sync"
	"sync/atomic"
	"time"

	"token-tracker/get"
	"token-tracker/logger"
	"token-tracker/utils"
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
