package infura

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"

	"github.com/spf13/viper"
)

// go test -v -run TestGetNodeClientVersion
func TestValidNodeClient(t *testing.T) {

	config.SetDevEnv()

	b, err := validNodeClient()
	if err != nil {
		t.Error(err)
	}

	if !b {
		t.Logf("Invalid Infura url: %s\n", viper.GetString("infura.https_endpoint"))
	} else {
		t.Logf("Valid Infura url: %s\n", viper.GetString("infura.https_endpoint"))
	}
}
