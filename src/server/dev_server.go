package server

import (
	"math"
	"time"

	"github.com/kyle-park-io/token-tracker/dev"
)

func StartDevServer() {
	go dev.ETLBlockData()

	time.Sleep(time.Duration(math.MaxInt64))
}
