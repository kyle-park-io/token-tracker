package get

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/request"
	"github.com/kyle-park-io/token-tracker/wss"
)

type EventLogsQuery struct {
	FromBlock string   `json:"fromBlock,omitempty"` // Starting block for the logs query
	ToBlock   string   `json:"toBlock,omitempty"`   // Ending block for the logs query
	BlockHash string   `json:"blockHash,omitempty"` // Block hash for the logs query
	Address   string   `json:"address,omitempty"`   // Address to filter logs by (contract or account)
	Topics    []string `json:"topics,omitempty"`    // Topics to filter specific events
}

func GetLogs(params EventLogsQuery) (interface{}, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getLogs", // Fetch logs from blockchain events, allowing event monitoring or analysis.
		Params:  []interface{}{params},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return "", err
	}

	emptyArray := []byte("[]")
	if bytes.Equal(resp.Result, emptyArray) {
		logger.Log.Infoln("GetLogs result is an empty array")
		// ws.GlobalLogChannel <- "GetLogs result is an empty array"
		wss.GlobalLogChannel <- "GetLogs result is an empty array"

		return "", nil
	}

	var eventLogs []map[string]interface{}
	if err := json.Unmarshal(resp.Result, &eventLogs); err != nil {
		return "", fmt.Errorf("failed to parse Result as EventLog: %w", err)
	}

	for i, eventLog := range eventLogs {
		_ = i
		for p, v := range eventLog {
			_ = p
			_ = v
		}
	}

	return eventLogs, nil
}
