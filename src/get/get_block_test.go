package get

import (
	"fmt"
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestGetBlockNumber
func TestGetBlockNumber(t *testing.T) {

	config.SetDevEnv()

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

	config.SetDevEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := false
	block, err := GetBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := viper.GetString("ROOT_PATH") + "/json/test/block/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockWithTxsByNumber
func TestGetBlockWithTxsByNumber(t *testing.T) {

	config.SetDevEnv()

	blockNumber, err := GetBlockNumber()
	if err != nil {
		t.Error(err)
	}

	withTxs := true
	block, err := GetBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	filePath := viper.GetString("ROOT_PATH") + "/json/test/block/block.json"
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockByNumber
func TestGetBlockByNumber(t *testing.T) {

	config.SetDevEnv()

	bn := int64(21491650)
	blockNumber := utils.DecimalToHex(bn)
	// blockNumber := "0x133ea62"
	withTxs := true

	block, err := GetBlockByNumber(string(blockNumber), withTxs)
	if err != nil {
		t.Error(err)
	}

	fileName := "block.json"
	folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/blocks/%s", blockNumber)
	filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/blocks/%s/%s", blockNumber, fileName)
	if err := utils.CreateFolderAndFile(folderPath, fileName); err != nil {
		t.Error(err)
	}
	if err := utils.SaveJSONToFile(block, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block. Check it in the JSON file.")
}

// go test -v -run TestGetBlockTimestampByNumber
func TestGetBlockTimestampByNumber(t *testing.T) {

	config.SetDevEnv()

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
