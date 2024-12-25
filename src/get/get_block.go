package get

import (
	"encoding/json"
	"fmt"

	"token-tracker/types/request"
	"token-tracker/types/response"
)

func GetBlockNumber() (response.BlockNumber, error) {
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

func GetBlockByNumber(blockNumber string, withTxs bool) (interface{}, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber", // Fetch the block for the requested block number
		Params:  []interface{}{blockNumber, withTxs},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return nil, err
	}

	// var block map[string]interface{}
	if !withTxs {
		var block response.BlockWithoutTransactions
		if err := json.Unmarshal(resp.Result, &block); err != nil {
			return "", fmt.Errorf("failed to parse Result as Block: %w", err)
		}
		return block, nil
	} else {
		var block response.BlockWithTransactions
		if err := json.Unmarshal(resp.Result, &block); err != nil {
			return "", fmt.Errorf("failed to parse Result as Block: %w", err)
		}
		return block, nil
	}
}

func GetBlockTimestampByNumber(blockNumber string) (string, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber", // Fetch the block for the requested block number
		Params:  []interface{}{blockNumber, false},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return "", err
	}

	var block response.BlockWithoutTransactions
	if err := json.Unmarshal(resp.Result, &block); err != nil {
		return "", fmt.Errorf("failed to parse Result as Block: %w", err)
	}

	return block.Timestamp, nil
}
