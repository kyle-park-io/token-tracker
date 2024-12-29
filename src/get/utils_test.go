package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/get/call"
	"github.com/kyle-park-io/token-tracker/internal/config"
)

// go test -v -run TestDecodeTransferLog
func TestDecodeTransferLog(t *testing.T) {

	// transfer event signature
	eventSignature := "Transfer(address,address,uint256)"
	eventHash := "0x" + call.Keccak256ToString([]byte(eventSignature))

	// transfer event data
	topics := []string{
		"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		"0x00000000000000000000000002f08670c84783473521850f37c3544c802b75c8",
		"0x00000000000000000000000014c53b68c8ac7bb5c0799e0ffd1abc72259cb506",
	}
	data := "0x0000000000000000000000000000000000000000000002d73297e0a2c9e0b000"

	if topics[0] != eventHash {
		t.Error("Event does not match Transfer signature")
		return
	}

	event := DecodeTransferLog(eventHash, topics, data)
	t.Log("Transfer Event Detected:")
	t.Logf("From: %s\n", event.From)
	t.Logf("To: %s\n", event.To)
	t.Logf("Value: %s\n", event.Value)
}

// go test -v -run TestDecodeRealTransferLog
func TestDecodeRealTransferLog(t *testing.T) {

	config.SetDevEnv()

	// Wrapped Ether Address
	address := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	params := EventLogsQuery{Address: address}
	e, err := GetLogs(params)
	if err != nil {
		t.Error(err)
	}

	// transfer event signature
	eventSignature := "Transfer(address,address,uint256)"
	eventHash := "0x" + call.Keccak256ToString([]byte(eventSignature))

	eventLogs, ok := e.([]map[string]interface{})
	if !ok {
		t.Error("check type.")
	}

	for _, eventLog := range eventLogs {
		rawTopics, ok := eventLog["topics"].([]interface{})
		if !ok {
			t.Error("check type.")
		}

		if rawTopics[0] == eventHash {

			topics := make([]string, len(rawTopics))
			for i, topic := range rawTopics {
				str, ok := topic.(string)
				if !ok {
					t.Error("expected string in topics")
					return
				}
				topics[i] = str
			}
			data, ok := eventLog["data"].(string)
			if !ok {
				t.Error("check type.")
			}

			event := DecodeTransferLog(eventHash, topics, data)
			t.Log("Transfer Event Detected:")
			t.Logf("From: %s\n", event.From)
			t.Logf("To: %s\n", event.To)
			t.Logf("Value: %s\n", event.Value)
		}
	}
}
