package get

import (
	"os"
	"testing"

	"token-tracker/configs"
	"token-tracker/logger"
	"token-tracker/utils"
)

// go test -v -run TestGetBlockNumber
func TestGetBlockNumber(t *testing.T) {

	setEnv()

	blockNumber, err := getBlockNumber()
	if err != nil {
		t.Error(err)
	}

	d, err := utils.HexToDecimal(string(blockNumber))
	if err != nil {
		t.Error(err)
	}

	t.Logf("The latest block number is: %s(%s)\n", blockNumber, d)
}

func setEnv() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
	if err := configs.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
