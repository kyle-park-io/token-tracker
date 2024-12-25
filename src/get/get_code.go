package get

import (
	"encoding/json"
	"fmt"

	"token-tracker/types/request"
)

// block tag
// latest, pending, earliest, 0x5BAD55
func getCode(address string) (string, error) {
	// JSON-RPC request data
	requestData := request.JSONRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getCode", // Fetch the bytecode of a smart contract at a specific address to check its existence or analyze its content.
		Params:  []interface{}{address, "latest"},
		ID:      1,
	}

	// Send the HTTP request
	resp, err := requestData.SendRequest()
	if err != nil {
		return "", err
	}

	var code string
	if err := json.Unmarshal(resp.Result, &code); err != nil {
		return "", fmt.Errorf("failed to parse Result as Code: %w", err)
	}

	return code, nil
}
