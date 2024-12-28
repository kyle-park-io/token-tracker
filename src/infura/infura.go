package infura

import (
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/types/request"
)

func validNodeClient() (bool, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "web3_clientVersion", // Fetch version of node client
		Params:  []interface{}{},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return false, err

	}

	var node string
	if err := json.Unmarshal(resp.Result, &node); err != nil {
		return false, fmt.Errorf("failed to parse Result as Node: %w", err)
	}
	if node == "" {
		return false, nil
	}

	return true, nil
}
