package router

import (
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/get/call"
	"github.com/kyle-park-io/token-tracker/integrated"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TrackERC20 godoc
// @Summary Track ERC20 token transfers
// @Description Tracks ERC20 token transfers for a specific account and token contract within a given date and target count.
// @Tags Track
// @Produce json
// @Param account query string true "Ethereum account address (e.g., '0x1234567890abcdef1234567890abcdef12345678')"
// @Param tokenAddress query string true "ERC20 token contract address (e.g., '0xabc123')"
// @Param date query string true "Date in 'YYYY-MM-DD' format"
// @Param targetCount query int false "Number of transactions to retrieve"
// @Param timeLimit query int false "Time limit for processing in seconds"
// @Success 200 {object} integrated.Result
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /track/trackERC20 [get]
func TrackERC20(c *gin.Context) {

	account := c.Query("account")
	// Ethereum address regex pattern
	pattern := `^0x[0-9a-fA-F]{40}$`
	matched, _ := regexp.MatchString(pattern, account)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'account' must be a valid Ethereum address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
		})
		return
	}
	lowercaseAccount := strings.ToLower(account)

	tokenAddress := c.Query("tokenAddress")
	matched, _ = regexp.MatchString(pattern, tokenAddress)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'tokenAddress' must be a valid Ethereum address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
		})
		return
	}

	inputCount := c.Query("targetCount")
	targetCount, err := strconv.Atoi(inputCount)
	if err != nil || targetCount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'targetCount' must be a positive integer",
		})
		return
	}

	date := c.Query("date")
	// Check if the input matches the format "YYYY-MM-DD"
	const layout = "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'date' must be in 'YYYY-MM-DD' format (e.g., '2024-01-01')",
		})
		return
	}
	yearMonthDay := date
	timezone := "UTC"

	timeLimit := c.Query("timeLimit")
	seconds, err := strconv.Atoi(timeLimit)
	if err != nil || seconds <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'timeLimit' must be a positive integer",
		})
		return
	}
	// Record the current time
	start := time.Now()

	day, _ := utils.ConvertYearMonthDayToTimezone(yearMonthDay, timezone)
	dayUnix, _ := utils.TimeToUnix(day, timezone)
	dayBigInt := big.NewInt(dayUnix)
	logger.Log.Infof("Day Unix: %d\n", dayUnix)

	nextDay, _ := utils.ConvertToNextDay(yearMonthDay, timezone)
	nextDayTime, _ := utils.ConvertYearMonthDayToTimezone(nextDay, timezone)
	nextDayUnix, _ := utils.TimeToUnix(nextDayTime, timezone)
	nextDayBigInt := big.NewInt(nextDayUnix)
	logger.Log.Infof("Next Day Unix: %d\n", nextDayUnix)

	blockPositionByDate, err := integrated.GetBlockPosition(yearMonthDay, timezone)
	if err != nil {
		logger.Log.Warnln(err)
	}
	logger.Log.Infof("%s Block Position: %+v\n", yearMonthDay, blockPositionByDate)

	blockPositionByNextDate, err := integrated.GetBlockPosition(nextDay, timezone)
	if err != nil {
		logger.Log.Warnln(err)
	}
	logger.Log.Infof("%s Block Position: %+v\n", nextDay, blockPositionByNextDate)

	// ERC20
	count := 0
	balance := ""
	balanceHex := ""
	transferHistory := make([]integrated.TransferHistory, 0)
	fromBlockNumber := ""
	toBlockNumber := ""

forLoop:
	for i := blockPositionByNextDate.High + 5; i >= blockPositionByDate.Low-5; i-- {

		// TODO: 429 Error
		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), false)
		if err != nil {
			logger.Log.Warnln(err)
		}
		block, _ := b.(response.BlockWithoutTransactions)

		bT := new(big.Int)
		blockTimestamp, _ := bT.SetString(block.Timestamp[2:], 16)
		if blockTimestamp.Cmp(nextDayBigInt) > 0 {
			continue
		} else {
			toBlockNumber = block.Number

			blockNumber, _ := utils.HexToDecimal(toBlockNumber)
			tag := utils.DecimalToHex(blockNumber - 1)

			// Example: ERC-20 balanceOf(address)
			methodName := "balanceOf"
			paramTypes := []string{"address"}
			params := []interface{}{account}

			callData, _ := call.CreateCallData(methodName, paramTypes, params)

			txArgs := map[string]interface{}{"to": tokenAddress, "data": callData}
			// TODO: 429 Error
			bal, err := get.GetCallBalance(txArgs, tag)
			if err != nil {
				logger.Log.Warnln(err)
			}
			balanceHex = utils.TrimLeadingZerosWithPrefix(string(bal))
			balance, _ = utils.HexToDecimalString(balanceHex)

			break forLoop
		}
	}

forLoop2:
	for i := blockPositionByDate.Low - 5; i <= blockPositionByNextDate.High+5; i++ {

		// TODO: 429 Error
		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), false)
		if err != nil {
			logger.Log.Warnln(err)
		}
		block, _ := b.(response.BlockWithoutTransactions)

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
	blockRanges := integrated.SplitBlockRange(toBlockNumber, fromBlockNumber, maxResults)
	for _, r := range blockRanges {

		logger.Log.Infof("Find Transfer Event, FromBlock: %s, ToBlock: %s\n", r[0], r[1])

		// Calculate elapsed time
		elapsed := time.Since(start).Seconds()
		if elapsed > float64(seconds) {
			result := integrated.Result{Account: account, TokenAddress: tokenAddress, Balance: balance, TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			_ = utils.CreateFolderAndFile(folderPath, fileName)
			_ = utils.SaveJSONToFile(result, filePath)

			c.JSON(http.StatusOK, result)
			return
		}

		// GetLogs
		params := get.EventLogsQuery{Address: tokenAddress, FromBlock: r[0], ToBlock: r[1],
			Topics: []string{eventHash}}
		// TODO: 429 Error
		e, err := get.GetLogs(params)
		if err != nil {
			logger.Log.Warnln(err)
		}

		// TODO: 429 Error
		eventLogs, ok := e.([]map[string]interface{})
		if !ok {
			logger.Log.Warnln(err)
		}
		for _, eventLog := range eventLogs {

			// TODO: 429 Error
			rawTopics, ok := eventLog["topics"].([]interface{})
			if !ok {
				logger.Log.Warnln(err)
			}

			fromTopic, ok := rawTopics[1].(string)
			if !ok {
				logger.Log.Warnln("topics[1] is not a string")
			}
			toTopic, ok := rawTopics[2].(string)
			if !ok {
				logger.Log.Warnln("topics[2] is not a string")
			}
			if "0x"+fromTopic[26:] != lowercaseAccount && "0x"+toTopic[26:] != lowercaseAccount {
				continue
			}

			topics := make([]string, len(rawTopics))
			for i, topic := range rawTopics {
				str, ok := topic.(string)
				if !ok {
					logger.Log.Warnln("expected string in topics")
				}
				topics[i] = str
			}
			data, _ := eventLog["data"].(string)

			blockNumber, _ := eventLog["blockNumber"].(string)
			txHash, _ := eventLog["transactionHash"].(string)
			// TODO: 429 Error
			b, err := get.GetBlockByNumber(blockNumber, false)
			if err != nil {
				logger.Log.Warnln(err)
			}
			block, _ := b.(response.BlockWithoutTransactions)

			unixTimestamp, _ := utils.HexToUnix(block.Timestamp)

			event := get.DecodeTransferLog(eventHash, topics, data)
			value, _ := utils.HexToDecimalString(event.Value)
			transferHistory = append(transferHistory, integrated.TransferHistory{TxHash: txHash, From: event.From,
				To: event.To, Value: value, ValueHex: event.Value, Timestamp: unixTimestamp})
			logger.Log.Infof("Transfer Info: from: %s, to: %s, value: %s\n", event.From, event.To, event.Value)

			count++
			logger.Log.Info("Event Count: ", count)
		}

		if count >= targetCount {

			result := integrated.Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			_ = utils.CreateFolderAndFile(folderPath, fileName)
			_ = utils.SaveJSONToFile(result, filePath)

			c.JSON(http.StatusOK, result)
			return
		}

		time.Sleep(2 * time.Second)
	}
}
