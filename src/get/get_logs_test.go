package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestGetLogs
func TestGetLogs(t *testing.T) {

	config.SetDevEnv()

	// Tether USD Contract Address
	address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	params := EventLogsQuery{Address: address}
	eventLogs, err := GetLogs(params)
	if err != nil {
		t.Error(err)
	}

	filePath := viper.GetString("ROOT_PATH") + "/json/test/eventLogs/eventLogs.json"
	if err := utils.SaveJSONToFile(eventLogs, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the event logs. Check it in the JSON file.")
}
