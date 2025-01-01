package router

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/gin-gonic/gin"
)

// GetRandomBlock godoc
// @Summary Retrieve a random block
// @Description Fetches a random block from the blockchain with optional transaction details.
// @Tags Block
// @Produce json
// @Param withTxs query boolean false "Include transactions in the block (true or false)"
// @Success 200 {object} response.BlockWithTransactions
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /get/getRandomBlock [get]
func GetRandomBlock(c *gin.Context) {

	// Retrieve the "withTxs" query parameter
	boolValue := c.Query("withTxs")

	// Attempt to parse the query parameter as a boolean
	withTxs, err := strconv.ParseBool(boolValue)
	if err != nil {
		// Respond with an error if parsing fails
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'withTxs' must be a boolean (true or false)",
		})
		return
	}

	// Get a random block number from the blockchain
	randomBlockNumber, err := get.GetRandomBlockNumber()
	if err != nil {
		// Log a warning if retrieving the block number fails
		logger.Log.Warnln(err)
	}

	// Fetch the block details using the random block number and the withTxs flag
	b, err := get.GetBlockByNumber(randomBlockNumber, withTxs)
	if err != nil {
		// Log a warning if fetching block details fails
		logger.Log.Warnln(err)
	}

	response := b
	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	// Send the block data as a JSON response
	c.JSON(http.StatusOK, response)
}

// GetBlock godoc
// @Summary Retrieve a block by number
// @Description Fetches a block from the blockchain based on its number with optional transaction details.
// @Tags Block
// @Produce json
// @Param number query string true "Block number in decimal or hexadecimal format"
// @Param withTxs query boolean false "Include transactions in the block (true or false)"
// @Success 200 {object} response.BlockWithTransactions
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /get/getBlock [get]
func GetBlock(c *gin.Context) {

	// Retrieve the "number" query parameter
	number := c.Query("number")
	blockNumber := ""

	// Check if the input is a decimal number
	if n, err := strconv.Atoi(number); err == nil {
		// Convert the decimal number to a hexadecimal string
		blockNumber = utils.DecimalToHex(int64(n))
	} else {
		// Check if the input is a valid hexadecimal number
		if strings.HasPrefix(number, "0x") {
			bigInt := new(big.Int)
			if _, success := bigInt.SetString(number[2:], 16); success {
				// If valid, use the hexadecimal input as is
				blockNumber = number
			} else {
				// Respond with an error if the input is invalid
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "'number' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
				})
				return
			}
		} else {
			// Return an error if the input is not a valid hexadecimal number
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'number' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
			})
			return
		}
	}

	// Retrieve the "withTxs" query parameter
	boolValue := c.Query("withTxs")

	// Attempt to parse the query parameter as a boolean
	withTxs, err := strconv.ParseBool(boolValue)
	if err != nil {
		// Respond with an error if parsing fails
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'withTxs' must be a boolean (true or false)",
		})
		return
	}

	// Fetch the block details using the block number and the withTxs flag
	b, err := get.GetBlockByNumber(blockNumber, withTxs)
	if err != nil {
		// Log a warning if fetching block details fails
		logger.Log.Warnln(err)
	}

	response := b
	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	// Send the block data as a JSON response
	c.JSON(http.StatusOK, response)
}
