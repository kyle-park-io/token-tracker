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

// TrackETH2 godoc
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
// @Router /track/trackETH2 [get]
func TrackETHBatch2(c *gin.Context) {

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
	balance := ""
	balanceHex := ""
	// transferHistory := make([]integrated.TransferHistory, 0)
	s := SafeSlice{}
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
			tag := utils.DecimalToHex(blockNumber)
			bal, err := get.GetBalance(account, tag)
			// TODO: 429 Error
			if err != nil {
				logger.Log.Warnln(err)
			}
			balanceHex = string(bal)
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

	maxResults := 50
	blockRanges := integrated.SplitBlockRange(toBlockNumber, fromBlockNumber, int64(maxResults))
	for _, r := range blockRanges {

		logger.Log.Infof("Find Transfer Event, FromBlock: %s, ToBlock: %s\n", r[0], r[1])
		// ws.GlobalLogChannel <- fmt.Sprintf("Find Transfer Event, FromBlock: %s, ToBlock: %s\n", r[0], r[1])
		wss.GlobalLogChannel <- fmt.Sprintf("Find Transfer Event, FromBlock: %s, ToBlock: %s\n", r[0], r[1])

		// Calculate elapsed time
		elapsed := time.Since(start).Seconds()
		if elapsed > float64(seconds) {
			transferHistory := s.GetAll()
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

		f, _ := utils.HexToDecimal(r[0])
		t, _ := utils.HexToDecimal(r[1])
		var r []get.RequestBalance
		for i := f; i <= t; i++ {
			r = append(r, get.RequestBalance{Address: account, Tag: utils.DecimalToHex(i)})
		}

		balances, err := get.GetBalanceBulk(r)
		if err != nil {
			logger.Log.Warnln(err)
		}
		bb, err := get.GetBalance(account, utils.DecimalToHex(t+1))
		if err != nil {
			logger.Log.Warnln(err)
		}

		var blockList []int
		for i := 0; i <= int(t)-int(f)-1; i++ {
			if balances[i] != balances[i+1] {
				blockList = append(blockList, i+int(f))
			}
		}
		if balances[len(balances)-1] != bb {
			blockList = append(blockList, int(t))
		}

		var r2 []string
		for i := 0; i < len(blockList); i++ {
			r2 = append(r2, utils.DecimalToHex(int64(blockList[i])))
		}
		// TODO: 429 Error
		b, err := get.GetBulkBlock(r2, true)
		if err != nil {
			logger.Log.Warnln(err)
		}

		go func() {

			// TODO: 429 Error
			blocks, _ := b.([]response.BlockWithTransactions)
			for _, block := range blocks {
				for _, transaction := range block.Transactions {
					if transaction.From != lowercaseAccount && transaction.To != lowercaseAccount {
						continue
					}
					if transaction.Value == "0x0" {
						continue
					}

					unixTimestamp, err := utils.HexToUnix(block.Timestamp)
					if err != nil {
						logger.Log.Error(err)
					}

					value, _ := utils.HexToDecimalString(transaction.Value)

					s.Append(integrated.TransferHistory{TxHash: transaction.Hash, From: transaction.From,
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
		}()

		if count >= targetCount {
			transferHistory := s.GetAll()
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

		time.Sleep(2 * time.Second)

	}

	transferHistory := s.GetAll()
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
}
