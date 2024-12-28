package get

import (
	"fmt"
	"testing"

	"token-tracker/configs"
	"token-tracker/types/response"
	"token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestRetrieveFullBlockchainData
func TestRetrieveFullBlockchainData(t *testing.T) {

	configs.SetEnv()

	randomBlockNumber, err := GetRandomBlockNumber()
	if err != nil {
		t.Error(err)
	}

	b, err := GetBlockByNumber(randomBlockNumber, true)
	if err != nil {
		t.Error(err)
	}
	block, ok := b.(response.BlockWithTransactions)
	if !ok {
		t.Error("check type.")
	}

	blockHash := block.Hash
	txHash := ""

forLoop:
	for _, v := range block.Transactions {
		txR, err := GetTransactionReceiptByHash(v.Hash)
		if err != nil {
			t.Error(err)
		}

		txReceipt, ok := txR.(response.TransactionReceipt)
		if !ok {
			t.Error("check type.")
		}
		if len(txReceipt.Logs) != 0 {
			txHash = txReceipt.TransactionHash
			break forLoop
		}
	}
	if txHash == "" {
		t.Error("please retry.")
	}

	tx, err := GetTransactionByHash(txHash)
	if err != nil {
		t.Error(err)
	}

	txReceipt, err := GetTransactionReceiptByHash(txHash)
	if err != nil {
		t.Error(err)
	}

	// address := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	params := EventLogsQuery{BlockHash: blockHash}
	eventLogs, err := GetLogs(params)
	if err != nil {
		t.Error(err)
	}

	folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/fullData/%s", randomBlockNumber)
	if err := utils.CreateFolder(folderPath); err != nil {
		t.Error(err)
	}

	blockName := "block.json"
	blockPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/fullData/%s/%s", randomBlockNumber, blockName)
	if err := utils.SaveJSONToFile(block, blockPath); err != nil {
		t.Error(err)
	}
	txName := "transaction.json"
	txPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/fullData/%s/%s", randomBlockNumber, txName)
	if err := utils.SaveJSONToFile(tx, txPath); err != nil {
		t.Error(err)
	}
	receiptName := "transactionReceipt.json"
	receiptPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/fullData/%s/%s", randomBlockNumber, receiptName)
	if err := utils.SaveJSONToFile(txReceipt, receiptPath); err != nil {
		t.Error(err)
	}
	logsName := "eventLogs.json"
	logsPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/fullData/%s/%s", randomBlockNumber, logsName)
	if err := utils.SaveJSONToFile(eventLogs, logsPath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the block full data. Check it in the JSON file.")
}
