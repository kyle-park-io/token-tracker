package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/configs"
	"github.com/kyle-park-io/token-tracker/get/call"
	"github.com/kyle-park-io/token-tracker/utils"
)

// go test -v -run TestGetCallBalance
func TestGetCallBalance(t *testing.T) {

	configs.SetEnv()

	var tag string
	tag = "latest"
	// Wrapped Ether Address (To)
	address := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

	// Example: ERC-20 balanceOf(address)
	methodName := "balanceOf"
	paramTypes := []string{"address"}
	params := []interface{}{"0x5c7BCd6E7De5423a257D81B442095A1a6ced35C5"}

	randomBlockNumber, err := GetRandomBlockNumber()
	if err != nil {
		t.Error(err)
	}
	if tag == "" {
		tag = randomBlockNumber
	}
	t.Logf("Random block: %s\n", randomBlockNumber)

	callData, err := call.CreateCallData(methodName, paramTypes, params)
	if err != nil {
		t.Error("Error:", err)
		return
	}

	txArgs := map[string]interface{}{"to": address, "data": callData}
	b, err := GetCallBalance(txArgs, tag)
	if err != nil {
		t.Error(err)
	}
	d, err := utils.HexToDecimalString(string(b))
	if err != nil {
		t.Error(err)
	}

	t.Logf("The account's balance is: %s(%s)\n", b, d)
}
