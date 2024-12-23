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

// go test -v -run TestGetBlockWithoutTxsByNumber
func TestGetBlockWithoutTxsByNumber(t *testing.T) {

	setEnv()

	blockNumber, err := getBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := false
	block, err := getBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockWithTxsByNumber
func TestGetBlockWithTxsByNumber(t *testing.T) {

	setEnv()

	blockNumber, err := getBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := true
	block, err := getBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

func setEnv() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
	if err := configs.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
