package dev

import (
	"fmt"
	"time"

	"github.com/kyle-park-io/token-tracker/db"
	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/response"
)

func ETLBlockData() {

	// db
	db.CreateDatabaseIfNotExists()
	db.InitDB()
	if err := db.DB.AutoMigrate(&db.Transaction{}, &db.TransactionReceipt{}, &db.Log{}); err != nil {
		fmt.Println(err)
		return
	}

	// data
	for {
		block, err := get.GetBulkBlockByRandomNumber(1, true)
		if err != nil {
			logger.Log.Error(err)
			time.Sleep(1 * time.Second)
			continue
		}

		switch v := block.(type) {
		case []response.BlockWithTransactions:
			for _, tx := range v[0].Transactions {
				txHash := tx.Hash

				tx, _ := get.GetTransactionByHash(txHash)
				if err := db.SaveBlockData(tx); err != nil {
					logger.Log.Error(err)
				}
				txR, _ := get.GetTransactionReceiptByHash(txHash)
				if err := db.SaveBlockData(txR); err != nil {
					logger.Log.Error(err)
				}
			}
		}
	}
}
