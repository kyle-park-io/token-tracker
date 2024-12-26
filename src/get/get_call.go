package get

import (
	"encoding/json"
	"fmt"

	"token-tracker/types/request"
	"token-tracker/types/response"
)

func GetCallBalance(txArgs map[string]interface{}, tag string) (response.Balance, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_call", // Fetch data from a smart contract using eth_call
		Params:  []interface{}{txArgs, tag},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		fmt.Print("hi")
		return "", err
	}

	// Extract the balance from the response result
	var balance response.Balance
	if err := json.Unmarshal(resp.Result, &balance); err != nil {
		return "", fmt.Errorf("failed to parse Result as Balance: %w", err)
	}

	return balance, nil
}
