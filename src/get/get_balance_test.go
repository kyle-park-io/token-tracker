package get

import (
	"testing"

	"token-tracker/configs"
	"token-tracker/utils"
)

// go test -v -run TestGetBalance
func TestGetBalance(t *testing.T) {

	configs.SetEnv()

	var tag string
	// Wrapped Ether Address
	address := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	// tag = "latest"

	randomBlockNumber, err := GetRandomBlockNumber()
	if err != nil {
		t.Error(err)
	}
	if tag == "" {
		tag = randomBlockNumber
	}
	t.Logf("Find block: %s\n", randomBlockNumber)

	b, err := GetBalance(address, tag)
	if err != nil {
		t.Error(err)
	}
	d, err := utils.HexToDecimalString(string(b))
	if err != nil {
		t.Error(err)
	}

	t.Logf("The account's balance is: %s(%s)\n", b, d)
}
