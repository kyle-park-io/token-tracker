package db

import (
	"fmt"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/spf13/viper"
)

func SaveBlockData(data interface{}) error {

	folderPath := viper.GetString("ROOT_PATH") + "/json/dev"
	txPath := viper.GetString("ROOT_PATH") + "/json/dev/transactions.json"
	txRPath := viper.GetString("ROOT_PATH") + "/json/dev/transactions-receipt.json"
	_ = utils.CreateFolderAndFile(folderPath, txPath)
	_ = utils.CreateFolderAndFile(folderPath, txRPath)

	switch v := data.(type) {
	case response.Transaction:
		_ = v

		j, _ := ReadInterfaceJSONFile(txPath)
		j = append(j, data)
		err := utils.SaveJSONToFile(j, txPath)
		if err != nil {
			logger.Log.Errorln(err)
		}
		return DB.Table("transactions").Create(data).Error
	case response.TransactionReceipt:
		_ = v

		j, _ := ReadInterfaceJSONFile(txRPath)
		j = append(j, data)
		err := utils.SaveJSONToFile(j, txRPath)
		if err != nil {
			logger.Log.Errorln(err)
		}
		_ = DB.Table("transaction_receipts").Create(data).Error
		return DB.Table("logs").Create(v.Logs).Error
	default:
		return fmt.Errorf("invalid data type: expected Transaction or TransacitonReceipt")
	}
}
