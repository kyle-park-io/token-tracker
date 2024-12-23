package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"token-tracker/types/response"

	"github.com/spf13/viper"
)

type JSONRPCRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	// Params interface{} `json:"params"`
	ID int `json:"id"`
}

func (j *JSONRPCRequest) SendRequest() (response.JSONRPCResponse, error) {
	// Encode the request data to JSON
	jsonData, err := json.Marshal(j)
	if err != nil {
		return response.JSONRPCResponse{}, fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	// Retrieve the Infura HTTPS endpoint from the configuration using Viper.
	infuraURL := viper.GetString("infura.https_endpoint")
	if infuraURL == "" {
		return response.JSONRPCResponse{}, fmt.Errorf("required configuration key %q is missing or empty", "infura.https_endpoint")
	}

	// Send an HTTP POST request
	resp, err := http.Post(infuraURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return response.JSONRPCResponse{}, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.JSONRPCResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response JSON
	var rpcResponse response.JSONRPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return response.JSONRPCResponse{}, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	// Handle RPC errors
	if rpcResponse.Error != nil {
		return response.JSONRPCResponse{}, fmt.Errorf("RPC error: %v", *rpcResponse.Error)
	}

	return rpcResponse, nil
}
