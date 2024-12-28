package get

import (
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/types/request"
	"github.com/kyle-park-io/token-tracker/types/response"
)

func GetBalance(address string, tag string) (response.Balance, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getBalance", // Fetch the account's balance
		Params:  []interface{}{address, tag},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return "", err
	}

	// Extract the balance from the response result
	var balance response.Balance
	if err := json.Unmarshal(resp.Result, &balance); err != nil {
		return "", fmt.Errorf("failed to parse Result as Balance: %w", err)
	}

	return balance, nil
}
