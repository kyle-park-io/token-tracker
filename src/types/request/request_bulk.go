package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/spf13/viper"
)

type JSONRPCRequestSlice []JSONRPCRequest

func (j JSONRPCRequestSlice) SendRequests() ([]response.JSONRPCResponse, error) {

	maxRetries := viper.GetInt("infura.eth.call_api_maxRetries")
	if maxRetries <= 0 {
		maxRetries = 100
	}
	retryInterval := viper.GetInt("infura.eth.retry_interval")
	if retryInterval <= 0 {
		retryInterval = 1
	}

	for attempt := 1; attempt <= maxRetries; attempt++ {

		// Encode the request data to JSON
		jsonData, err := json.Marshal(j)
		if err != nil {
			return []response.JSONRPCResponse{}, fmt.Errorf("failed to marshal JSON data: %v", err)
		}

		// Retrieve the Infura HTTPS endpoint from the configuration using Viper.
		infuraURL := viper.GetString("infura.https_endpoint")
		if infuraURL == "" {
			return []response.JSONRPCResponse{}, fmt.Errorf("required configuration key %q is missing or empty", "infura.https_endpoint")
		}

		// Send an HTTP POST request
		resp, err := http.Post(infuraURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return []response.JSONRPCResponse{}, fmt.Errorf("failed to send HTTP request: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []response.JSONRPCResponse{}, fmt.Errorf("failed to read response body: %v", err)
		}

		// Parse the response JSON
		var rpcResponse []response.JSONRPCResponse
		err = json.Unmarshal(body, &rpcResponse)
		if err != nil {
			return []response.JSONRPCResponse{}, fmt.Errorf("failed to unmarshal JSON response: %v", err)
		}

		// Handle RPC errors
		if rpcResponse[0].Error != nil {
			if isRateLimitError(rpcResponse[0].Error) {
				//TODO: 429 Error
				// ws.GlobalLogChannel <- fmt.Sprintf("Rate limit error (429) encountered. Retrying... Attempt %d/%d", attempt, maxRetries)
				wss.GlobalLogChannel <- fmt.Sprintf("Rate limit error (429) encountered. Retrying... Attempt %d/%d", attempt, maxRetries)
				time.Sleep(time.Duration(retryInterval) * time.Second)
			} else {
				e := fmt.Errorf("RPC error: %v", *rpcResponse[0].Error)
				// ws.GlobalLogChannel <- e.Error()
				wss.GlobalLogChannel <- e.Error()
				logger.Log.Warnln("Maybe here.")
				logger.Log.Warnf("%+v\n", j)
				time.Sleep(120 * time.Second)
				return []response.JSONRPCResponse{}, e
			}
		} else {
			return rpcResponse, nil
		}
	}

	finalErr := fmt.Errorf("RPC call failed after %d attempts", maxRetries)
	// ws.GlobalLogChannel <- finalErr.Error()
	wss.GlobalLogChannel <- finalErr.Error()
	return []response.JSONRPCResponse{}, finalErr
}
