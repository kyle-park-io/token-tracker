package get

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
)

// go test -v -run TestGetCodeForEOA
func TestGetCodeForEOA(t *testing.T) {

	config.SetDevEnv()

	// Kyle Test Address
	address := "0xF50fe83f158c4D330759694b9fC03D3C39779926"
	code, err := getCode(address)
	if err != nil {
		t.Error(err)
	}

	if code == "0x" {
		t.Logf("This address(%s) belongs to an externally owned account (EOA).\n", address)
	} else {
		t.Logf("This address(%s) belongs to a smart contract.\n", address)
	}
}

// go test -v -run TestGetCodeForContract
func TestGetCodeForContract(t *testing.T) {

	config.SetDevEnv()

	// Tether USD Contract Address
	address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	code, err := getCode(address)
	if err != nil {
		t.Error(err)
	}

	if code == "0x" {
		t.Logf("This address(%s) belongs to an externally owned account (EOA).\n", address)
	} else {
		t.Logf("This address(%s) belongs to a smart contract.\n", address)
	}
}
