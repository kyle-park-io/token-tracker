package get

import (
	"fmt"
	"testing"

	"token-tracker/configs"
	"token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestGetTransactionByHash
func TestGetTransactionByHash(t *testing.T) {

	configs.SetEnv()

	txHash := "0x797064084c35761950503758847f0655791880eab1df8d2983784dc20fc32391"
	transaction, err := GetTransactionByHash(txHash)
	if err != nil {
		t.Error(err)
	}

	fileName := "transaction.json"
	folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transactions/%s", txHash)
	filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transactions/%s/%s", txHash, fileName)
	if err := utils.CreateFolderAndFile(folderPath, fileName); err != nil {
		t.Error(err)
	}
	if err := utils.SaveJSONToFile(transaction, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the transaction. Check it in the JSON file.")
}

// go test -v -run TestGetTransactionReceiptByHash
func TestGetTransactionReceiptByHash(t *testing.T) {

	configs.SetEnv()

	txHash := "0x797064084c35761950503758847f0655791880eab1df8d2983784dc20fc32391"
	txReceipt, err := GetTransactionReceiptByHash(txHash)
	if err != nil {
		t.Error(err)
	}

	fileName := "transactionReceipt.json"
	folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transactions/%s", txHash)
	filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transactions/%s/%s", txHash, fileName)
	if err := utils.CreateFolderAndFile(folderPath, fileName); err != nil {
		t.Error(err)
	}
	if err := utils.SaveJSONToFile(txReceipt, filePath); err != nil {
		t.Error(err)
	}

	t.Logf("Successfully fetched the transaction receipt. Check it in the JSON file.")
}
