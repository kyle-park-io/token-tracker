package get

import (
	"encoding/json"
	"fmt"

	"token-tracker/types/request"
	"token-tracker/types/response"
)

func getBlockNumber() (response.BlockNumber, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber", // Fetch the latest block number
		Params:  []interface{}{},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return "", err
	}

	// Extract the block number from the response result
	var blockNumber response.BlockNumber
	if err := json.Unmarshal(resp.Result, &blockNumber); err != nil {
		return "", fmt.Errorf("failed to parse Result as BlockNumber: %w", err)
	}

	return blockNumber, nil
}
