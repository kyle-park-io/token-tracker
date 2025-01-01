package router

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/integrated"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/types/response"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TrackETH godoc
// @Summary Track Ethereum account balance
// @Description Tracks Ethereum account activity, including transactions and balance changes, within a given date and target count.
// @Tags Track
// @Produce json
// @Param account query string true "Ethereum account address (e.g., '0x1234567890abcdef1234567890abcdef12345678')"
// @Param date query string true "Date in 'YYYY-MM-DD' format"
// @Param targetCount query int false "Number of transactions to retrieve"
// @Param timeLimit query int false "Time limit for processing in seconds"
// @Success 200 {object} integrated.Result
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /track/trackETH [get]
func TrackETH(c *gin.Context) {

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

	tokenAddress := "ETH"

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
	// ws.GlobalLogChannel <- fmt.Sprintf("%s Block Position: %+v\n", yearMonthDay, blockPositionByDate)
	wss.GlobalLogChannel <- fmt.Sprintf("%s Block Position: %+v\n", yearMonthDay, blockPositionByDate)

	blockPositionByNextDate, err := integrated.GetBlockPosition(nextDay, timezone)
	if err != nil {
		logger.Log.Warnln(err)
	}
	logger.Log.Infof("%s Block Position: %+v\n", nextDay, blockPositionByNextDate)
	// ws.GlobalLogChannel <- fmt.Sprintf("%s Block Position: %+v\n", nextDay, blockPositionByNextDate)
	wss.GlobalLogChannel <- fmt.Sprintf("%s Block Position: %+v\n", nextDay, blockPositionByNextDate)

	// ETH
	count := 0
	blockCount := 0
	transferHistory := make([]integrated.TransferHistory, 0)
	for i := blockPositionByNextDate.High + 5; i >= blockPositionByDate.Low-5; i-- {

		// TODO: 429 Error
		b, err := get.GetBlockByNumber(utils.DecimalToHex(i), true)
		if err != nil {
			logger.Log.Warnln(err)
		}
		block, _ := b.(response.BlockWithTransactions)

		balance := ""
		balanceHex := ""
		bT := new(big.Int)
		blockTimestamp, _ := bT.SetString(block.Timestamp[2:], 16)
		if blockTimestamp.Cmp(nextDayBigInt) > 0 {
			continue
		} else if balance == "" {
			blockNumber, _ := utils.HexToDecimal(block.Number)
			tag := utils.DecimalToHex(blockNumber - 1)
			bal, err := get.GetBalance(account, tag)
			// TODO: 429 Error
			if err != nil {
				logger.Log.Warnln(err)
			}
			balanceHex = string(bal)
			balance, _ = utils.HexToDecimalString(balanceHex)
		}

		if blockTimestamp.Cmp(dayBigInt) < 0 {

			result := integrated.Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			_ = utils.CreateFolderAndFile(folderPath, fileName)
			_ = utils.SaveJSONToFile(result, filePath)

			jsonData, _ := json.Marshal(result)
			// ws.GlobalLogChannel <- string(jsonData)
			wss.GlobalLogChannel <- string(jsonData)

			c.JSON(http.StatusOK, result)
			return
		}

		// Calculate elapsed time
		elapsed := time.Since(start).Seconds()
		if elapsed > float64(seconds) {
			result := integrated.Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			_ = utils.CreateFolderAndFile(folderPath, fileName)
			_ = utils.SaveJSONToFile(result, filePath)

			jsonData, _ := json.Marshal(result)
			// ws.GlobalLogChannel <- string(jsonData)
			wss.GlobalLogChannel <- string(jsonData)

			c.JSON(http.StatusOK, result)
			return
		}

		for _, transaction := range block.Transactions {

			if transaction.From == lowercaseAccount || transaction.To == lowercaseAccount {
				if transaction.Value == "0x0" {
					continue
				}

				unixTimestamp, err := utils.HexToUnix(block.Timestamp)
				if err != nil {
					logger.Log.Error(err)
				}

				value, _ := utils.HexToDecimalString(transaction.Value)
				transferHistory = append(transferHistory, integrated.TransferHistory{TxHash: transaction.Hash, From: transaction.From,
					To: transaction.To, Value: value, ValueHex: transaction.Value, Timestamp: unixTimestamp})
				logger.Log.Infof("Transfer Info: from: %s, to: %s, value: %s\n", transaction.From, transaction.To, transaction.Value)
				// ws.GlobalLogChannel <- fmt.Sprintf("Transfer Info: from: %s, to: %s, value: %s\n", transaction.From, transaction.To, transaction.Value)
				wss.GlobalLogChannel <- fmt.Sprintf("Transfer Info: from: %s, to: %s, value: %s\n", transaction.From, transaction.To, transaction.Value)

				count++
				logger.Log.Info("Event Count: ", count)
				// ws.GlobalLogChannel <- fmt.Sprintf("Event Count: %d\n", count)
				wss.GlobalLogChannel <- fmt.Sprintf("Event Count: %d\n", count)
			}
		}
		blockCount++
		logger.Log.Info("Block Count: ", blockCount)
		// ws.GlobalLogChannel <- fmt.Sprintf("Block Count: %d\n", count)
		wss.GlobalLogChannel <- fmt.Sprintf("Block Count: %d\n", blockCount)

		if count >= targetCount {

			result := integrated.Result{Account: account, TokenAddress: tokenAddress, Balance: balance, BalanceHex: balanceHex,
				TransferHistory: transferHistory}
			fileName := tokenAddress + ".json"
			folderPath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s", account)
			filePath := viper.GetString("ROOT_PATH") + fmt.Sprintf("/json/transferHistory/%s/%s", account, fileName)
			_ = utils.CreateFolderAndFile(folderPath, fileName)
			_ = utils.SaveJSONToFile(result, filePath)

			jsonData, _ := json.Marshal(result)
			// ws.GlobalLogChannel <- string(jsonData)
			wss.GlobalLogChannel <- string(jsonData)

			c.JSON(http.StatusOK, result)
			return
		}
	}
}
