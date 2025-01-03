package integrated

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestTrackETH
func TestTrackETH(t *testing.T) {

	// Initialize configuration environment
	config.SetDevEnv()

	// Test input values
	// Wrapped Ether Address
	account := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	lowercaseAccount := strings.ToLower(account)
	tokenAddress := "ETH"
	targetCount := 10
	yearMonthDay := "2024-12-24"
	timezone := "UTC"

	day, err := utils.ConvertYearMonthDayToTimezone(yearMonthDay, timezone)
	if err != nil {
		t.Error(err)
	}
	dayUnix, err := utils.TimeToUnix(day, timezone)
	if err != nil {
		t.Error(err)
	}
	dayBigInt := big.NewInt(dayUnix)
	t.Logf("Day Unix: %d\n", dayUnix)

	nextDay, err := utils.ConvertToNextDay(yearMonthDay, timezone)
	if err != nil {
		t.Error(err)
	}
	nextDayTime, err := utils.ConvertYearMonthDayToTimezone(nextDay, timezone)
	if err != nil {
		t.Errorf("Failed to convert to next day: %v", err)
	}
	nextDayUnix, err := utils.TimeToUnix(nextDayTime, timezone)
	if err != nil {
		t.Error(err)
	}
	nextDayBigInt := big.NewInt(nextDayUnix)
	t.Logf("Next Day Unix: %d\n", nextDayUnix)

	blockPositionByDate, err := GetBlockPosition(yearMonthDay, timezone)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s Block Position: %+v\n", yearMonthDay, blockPositionByDate)

	blockPositionByNextDate, err := GetBlockPosition(nextDay, timezone)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s Block Position: %+v\n", nextDay, blockPositionByNextDate)

	// ETH
	count := 0
	blockCount := 0
	transferHistory := make([]TransferHistory, 0)
	for i := blockPositionByNextDate.High + 5; i >= blockPositionByDate.Low-5; i-- {

		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), true)
		if err != nil {
			// t.Error(err)
			t.Fatal(err)
		}

		block, ok := b.(response.BlockWithTransactions)
		if !ok {
			t.Error("check type.")
		}

		balance := ""
		balanceHex := ""
		bT := new(big.Int)
		blockTimestamp, _ := bT.SetString(block.Timestamp[2:], 16)
		if blockTimestamp.Cmp(nextDayBigInt) > 0 {
			continue
		} else if balance == "" {
			blockNumber, _ := utils.HexToDecimal(block.Number)
			tag := utils.DecimalToHex(blockNumber)
			bal, err := get.GetBalance(account, tag)
			if err != nil {
				t.Error(err)
			}
			balanceHex = string(bal)
			balance, _ = utils.HexToDecimalString(balanceHex)
		}

		if blockTimestamp.Cmp(dayBigInt) < 0 {

			result := Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			if err := utils.CreateFolderAndFile(folderPath, fileName); err != nil {
				t.Error(err)
			}
			if err := utils.SaveJSONToFile(result, filePath); err != nil {
				t.Error(err)
			}

			return
		}

		for _, transaction := range block.Transactions {

			if transaction.From == lowercaseAccount || transaction.To == lowercaseAccount {
				if transaction.Value == "0x0" {
					continue
				}

				unixTimestamp, err := utils.HexToUnix(block.Timestamp)
				if err != nil {
					t.Error(err)
				}

				value, _ := utils.HexToDecimalString(transaction.Value)
				transferHistory = append(transferHistory, TransferHistory{TxHash: transaction.Hash, From: transaction.From,
					To: transaction.To, Value: value, ValueHex: transaction.Value, Timestamp: unixTimestamp})
				t.Logf("Transfer Info: from: %s, to: %s, value: %s\n", transaction.From, transaction.To, transaction.Value)

				count++
				t.Log("Event Count: ", count)
			}
		}
		blockCount++
		t.Log("Block Count: ", blockCount)

		if count >= targetCount {

			result := Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			if err := utils.CreateFolderAndFile(folderPath, fileName); err != nil {
				t.Error(err)
			}
			if err := utils.SaveJSONToFile(result, filePath); err != nil {
				t.Error(err)
			}

			return
		}
	}
}
