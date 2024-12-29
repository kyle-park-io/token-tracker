package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/utils"
)

// go test -v -run TestGetBalance
func TestGetBalance(t *testing.T) {

	config.SetDevEnv()

	var tag string
	// tag = "latest"
	// Wrapped Ether Address
	address := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

	randomBlockNumber, err := GetRandomBlockNumber()
	if err != nil {
		t.Error(err)
	}
	if tag == "" {
		tag = randomBlockNumber
	}
	t.Logf("Random block: %s\n", randomBlockNumber)

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
