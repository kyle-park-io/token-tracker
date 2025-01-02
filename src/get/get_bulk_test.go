package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
)

// go test -v -run TestGetBulkBlockByRandomNumber
func TestGetBulkBlockByRandomNumber(t *testing.T) {

	config.SetDevEnv()

	result, err := GetBulkBlockByRandomNumber(10, false)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)
}
