package get

import (
	"encoding/json"
	"fmt"

	"github.com/kyle-park-io/token-tracker/types/request"
	"github.com/kyle-park-io/token-tracker/types/response"
)

func GetTransactionByHash(txHash string) (interface{}, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionByHash", // Fetch the transaction for the requested transaction hash
		Params:  []interface{}{txHash},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return nil, err
	}

	var transaction response.Transaction
	if err := json.Unmarshal(resp.Result, &transaction); err != nil {
		return "", fmt.Errorf("failed to parse Result as Transaction: %w", err)
	}

	return transaction, nil
}

func GetTransactionReceiptByHash(txHash string) (interface{}, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionReceipt", // Fetch the transaction receipt for the requested transaction hash
		Params:  []interface{}{txHash},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return nil, err
	}

	var transactionReceipt response.TransactionReceipt
	if err := json.Unmarshal(resp.Result, &transactionReceipt); err != nil {
		return "", fmt.Errorf("failed to parse Result as TransactionReceipt: %w", err)
	}

	return transactionReceipt, nil
}
