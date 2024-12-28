package tracker

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/configs"
)

// go test -v -run TestTrackBlockTimestamp
func TestTrackBlockTimestamp(t *testing.T) {

	configs.SetEnv()

	blockTimestamp := int64(1672113600)
	position, err := TrackBlockTimestamp(blockTimestamp)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Block Position: %+v", position)
}
