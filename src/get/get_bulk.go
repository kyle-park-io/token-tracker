package get

import (
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/types/request"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"
)

func GetBulkBlockByRandomNumber(blockCount int, withTxs bool) (interface{}, error) {

	blockNumber, err := GetBlockNumber()
	if err != nil {
		return "", err
	}

	var requests request.JSONRPCRequestSlice
	for i := 0; i < blockCount; i++ {
		r, err := utils.RandomHexBelow(string(blockNumber))
		if err != nil {
			return "", err
		}

		// JSON-RPC request data
		requestData := request.JSONRPCRequest{
			JsonRpc: "2.0",
			Method:  "eth_getBlockByNumber", // Fetch the block for the requested block number
			Params:  []interface{}{r, withTxs},
			ID:      1,
		}
		requests = append(requests, requestData)
	}

	// Send the HTTP request
	resp, err := requests.SendRequests()
	if err != nil {
		return nil, err
	}

	// var block map[string]interface{}
	if !withTxs {
		var blocks []response.BlockWithoutTransactions
		for _, v := range resp {

			var block response.BlockWithoutTransactions
			if err := json.Unmarshal(v.Result, &block); err != nil {
				return "", fmt.Errorf("failed to parse Result as Block: %w", err)
			}
			blocks = append(blocks, block)
		}
		return blocks, nil
	} else {

		var blocks []response.BlockWithTransactions
		for _, v := range resp {
			var block response.BlockWithTransactions
			if err := json.Unmarshal(v.Result, &block); err != nil {
				return "", fmt.Errorf("failed to parse Result as Block: %w", err)
			}
			blocks = append(blocks, block)
		}
		return blocks, nil
	}
}

func GetBulkBlock(r []string, withTxs bool) (interface{}, error) {

	var requests request.JSONRPCRequestSlice
	for _, v := range r {
		// JSON-RPC request data
		requestData := request.JSONRPCRequest{
			JsonRpc: "2.0",
			Method:  "eth_getBlockByNumber", // Fetch the block for the requested block number
			Params:  []interface{}{v, withTxs},
			ID:      1,
		}
		requests = append(requests, requestData)
	}

	// Send the HTTP request
	resp, err := requests.SendRequests()
	if err != nil {
		return nil, err
	}

	// var block map[string]interface{}
	if !withTxs {
		var blocks []response.BlockWithoutTransactions
		for _, v := range resp {

			var block response.BlockWithoutTransactions
			if err := json.Unmarshal(v.Result, &block); err != nil {
				return "", fmt.Errorf("failed to parse Result as Block: %w", err)
			}
			blocks = append(blocks, block)
		}
		return blocks, nil
	} else {

		var blocks []response.BlockWithTransactions
		for _, v := range resp {
			var block response.BlockWithTransactions
			if err := json.Unmarshal(v.Result, &block); err != nil {
				return "", fmt.Errorf("failed to parse Result as Block: %w", err)
			}
			blocks = append(blocks, block)
		}
		return blocks, nil
	}
}
