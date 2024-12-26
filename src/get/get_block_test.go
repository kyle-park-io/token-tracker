package get

import (
	"testing"

	"token-tracker/configs"
	"token-tracker/utils"
)

// go test -v -run TestGetBlockNumber
func TestGetBlockNumber(t *testing.T) {

	configs.SetEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	d, err := utils.HexToDecimalString(string(blockNumber))
	if err != nil {
		t.Error(err)
	}

	t.Logf("The latest block number is: %s(%s)\n", blockNumber, d)
}

// go test -v -run TestGetBlockWithoutTxsByNumber
func TestGetBlockWithoutTxsByNumber(t *testing.T) {

	configs.SetEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := false
	block, err := GetBlockByNumber(string(blockNumber), withTxs)
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

	configs.SetEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := true
	block, err := GetBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockByNumber
func TestGetBlockByNumber(t *testing.T) {

	configs.SetEnv()

	blockNumber := "0x133ea62"
	withTxs := true

	block, err := GetBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockTimestampByNumber
func TestGetBlockTimestampByNumber(t *testing.T) {

	configs.SetEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	timestamp, err := GetBlockTimestampByNumber(string(blockNumber))
	if err != nil {
		t.Error(err)
	}

	t.Log("Hex Timestamp:", timestamp)
}
