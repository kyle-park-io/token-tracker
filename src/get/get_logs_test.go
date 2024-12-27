package get

import (
	"testing"

	"token-tracker/configs"
	"token-tracker/utils"
)

// go test -v -run TestGetLogs
func TestGetLogs(t *testing.T) {

	configs.SetEnv()

	// Tether USD Contract Address
	address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	params := EventLogsQuery{Address: address}
	eventLogs, err := GetLogs(params)
	if err != nil {
		t.Error(err)
	}

	filePath := "/home/kyle/code/token-tracker/src/json/eventLogs/eventLogs.json"
	if err := utils.SaveJSONToFile(eventLogs, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the event logs. Check it in the JSON file.")
}
