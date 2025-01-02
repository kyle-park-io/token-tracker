package get

import (
	"fmt"
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

// go test -v -run TestGetBalanceBulk
func TestGetBalanceBulk(t *testing.T) {

	config.SetDevEnv()

	// Wrapped Ether Address
	address := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	var requests []RequestBalance
	for i := 0; i < 10; i++ {
		randomBlockNumber, err := GetRandomBlockNumber()
		if err != nil {
			t.Error(err)
		}
		request := RequestBalance{Address: address, Tag: randomBlockNumber}
		requests = append(requests, request)
	}

	balances, err := GetBalanceBulk(requests)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(balances)

	for _, v := range balances {
		d, err := utils.HexToDecimalString(string(v))
		if err != nil {
			t.Error(err)
		}

		t.Logf("The account's balance is: %s(%s)\n", v, d)
	}
}
