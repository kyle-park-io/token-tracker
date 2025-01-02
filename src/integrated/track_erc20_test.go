package integrated

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/get/call"
	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/spf13/viper"
)

// go test -v -run TestTrackERC20
func TestTrackERC20(t *testing.T) {

	// Initialize configuration environment
	config.SetDevEnv()

	// Test input values
	account := "0x1E2aB9200B2Fe8832A55CACCB917872dB2715C31"
	lowercaseAccount := strings.ToLower(account)
	// Wrapped Ether Address
	tokenAddress := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	targetCount := 1
	yearMonthDay := "2024-12-16"
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

	// ERC20
	count := 0
	balance := ""
	balanceHex := ""
	transferHistory := make([]TransferHistory, 0)
	fromBlockNumber := ""
	toBlockNumber := ""

forLoop:
	for i := blockPositionByNextDate.High + 5; i >= blockPositionByDate.Low-5; i-- {

		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), false)
		if err != nil {
			// t.Error(err)
			t.Fatal(err)
		}

		block, ok := b.(response.BlockWithoutTransactions)
		if !ok {
			t.Error("check type.")
		}

		bT := new(big.Int)
		blockTimestamp, _ := bT.SetString(block.Timestamp[2:], 16)
		if blockTimestamp.Cmp(nextDayBigInt) > 0 {
			continue
		} else {
			toBlockNumber = block.Number

			blockNumber, _ := utils.HexToDecimal(toBlockNumber)
			tag := utils.DecimalToHex(blockNumber)

			// Example: ERC-20 balanceOf(address)
			methodName := "balanceOf"
			paramTypes := []string{"address"}
			params := []interface{}{account}

			callData, err := call.CreateCallData(methodName, paramTypes, params)
			if err != nil {
				t.Error("Error:", err)
				return
			}

			txArgs := map[string]interface{}{"to": tokenAddress, "data": callData}
			bal, err := get.GetCallBalance(txArgs, tag)
			if err != nil {
				t.Error(err)
			}
			balanceHex = utils.TrimLeadingZerosWithPrefix(string(bal))
			balance, _ = utils.HexToDecimalString(balanceHex)

			break forLoop
		}
	}

forLoop2:
	for i := blockPositionByDate.Low - 5; i <= blockPositionByNextDate.High+5; i++ {

		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), false)
		if err != nil {
			// t.Error(err)
			t.Fatal(err)
		}

		block, ok := b.(response.BlockWithoutTransactions)
		if !ok {
			t.Error("check type.")
		}

		bT := new(big.Int)
		blockTimestamp, _ := bT.SetString(block.Timestamp[2:], 16)

		if blockTimestamp.Cmp(dayBigInt) >= 0 {
			fromBlockNumber = block.Number
			break forLoop2
		}
	}

	// transfer event signature
	eventSignature := "Transfer(address,address,uint256)"
	eventHash := "0x" + call.Keccak256ToString([]byte(eventSignature))

	maxResults := int64(200)
	blockRanges := SplitBlockRange(toBlockNumber, fromBlockNumber, maxResults)
	for _, r := range blockRanges {
		t.Logf("Find Transfer Event, FromBlock: %s, ToBlock: %s\n", r[0], r[1])

		// GetLogs
		params := get.EventLogsQuery{Address: tokenAddress, FromBlock: r[0], ToBlock: r[1],
			Topics: []string{eventHash}}
		e, err := get.GetLogs(params)
		if err != nil {
			t.Error(err)
		}

		eventLogs, ok := e.([]map[string]interface{})
		if !ok {
			t.Error("check type.")
		}

		for _, eventLog := range eventLogs {

			rawTopics, ok := eventLog["topics"].([]interface{})
			if !ok {
				t.Error("check type.")
			}

			fromTopic, ok := rawTopics[1].(string)
			if !ok {
				t.Error("topics[1] is not a string")
				return
			}
			toTopic, ok := rawTopics[2].(string)
			if !ok {
				t.Error("topics[2] is not a string")
				return
			}
			if "0x"+fromTopic[26:] != lowercaseAccount && "0x"+toTopic[26:] != lowercaseAccount {
				continue
			}

			topics := make([]string, len(rawTopics))
			for i, topic := range rawTopics {
				str, ok := topic.(string)
				if !ok {
					t.Error("expected string in topics")
					return
				}
				topics[i] = str
			}
			data, ok := eventLog["data"].(string)
			if !ok {
				t.Error("check type.")
			}

			blockNumber, _ := eventLog["blockNumber"].(string)
			txHash, _ := eventLog["transactionHash"].(string)
			b, err := get.GetBlockByNumber(blockNumber, false)
			if err != nil {
				t.Error(err)
			}
			block, ok := b.(response.BlockWithoutTransactions)
			if !ok {
				t.Error("check type.")
			}
			unixTimestamp, _ := utils.HexToUnix(block.Timestamp)

			event := get.DecodeTransferLog(eventHash, topics, data)
			value, _ := utils.HexToDecimalString(event.Value)
			transferHistory = append(transferHistory, TransferHistory{TxHash: txHash, From: event.From,
				To: event.To, Value: value, ValueHex: event.Value, Timestamp: unixTimestamp})
			t.Logf("Transfer Info: from: %s, to: %s, value: %s\n", event.From, event.To, event.Value)

			count++
			t.Log("Event Count: ", count)
		}

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
		}

		time.Sleep(2 * time.Second)
	}

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
}
