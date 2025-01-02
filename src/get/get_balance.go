package get

import (
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/types/request"
	"github.com/kyle-park-io/token-tracker/types/response"
)

type RequestBalance struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

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

func GetBalanceBulk(r []RequestBalance) ([]response.Balance, error) {

	var requests request.JSONRPCRequestSlice
	for _, v := range r {
		// JSON-RPC request data
		requestData := request.JSONRPCRequest{
			JsonRpc: "2.0",
			Method:  "eth_getBalance", // Fetch the account's balance
			Params:  []interface{}{v.Address, v.Tag},
			ID:      1,
		}
		requests = append(requests, requestData)
	}

	// Send the HTTP request
	resp, err := requests.SendRequests()
	if err != nil {
		return nil, err
	}

	// Extract the balance from the response result
	var balances []response.Balance
	for _, v := range resp {
		var balance response.Balance
		if err := json.Unmarshal(v.Result, &balance); err != nil {
			return nil, fmt.Errorf("failed to parse Result as Balance: %w", err)
		}
		balances = append(balances, balance)
	}

	return balances, nil
}
