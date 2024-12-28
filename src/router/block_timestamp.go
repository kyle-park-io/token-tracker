package router

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"

	"github.com/gin-gonic/gin"
)

type ResponseBlockTimestamp struct {
	Timestamp    int64     `json:"timestamp" example:"1672531199"`    // Unix timestamp
	HexTimestamp string    `json:"hexTimestamp" example:"0x62d43b1f"` // Hexadecimal timestamp
	Date         time.Time `json:"date" example:"2024-01-01"`         // Original date
}

// GetBlockTimestamp godoc
// @Summary Retrieve block timestamp by block number
// @Description Fetches the block timestamp for a specific block number in both Unix and hexadecimal formats.
// @Tags Block
// @Produce json
// @Param number query string true "Block number in decimal or hexadecimal format"
// @Success 200 {object} ResponseBlockTimestamp
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /get/getBlockTimestamp [get]
func GetBlockTimestamp(c *gin.Context) {

	number := c.Query("number")
	blockNumber := ""
	// Check if the input is a decimal number
	if n, err := strconv.Atoi(number); err == nil {
		blockNumber = utils.DecimalToHex(int64(n))
	}
	// Check if the input is a hexadecimal number
	if strings.HasPrefix(number, "0x") {
		bigInt := new(big.Int)
		if _, success := bigInt.SetString(number[2:], 16); success {
			blockNumber = number
		} else {
			// Return an error if the input is not a valid hexadecimal number
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'number' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
			})
			return
		}
	}

	hexTimestamp, err := get.GetBlockTimestampByNumber(blockNumber)
	if err != nil {
		logger.Log.Warnln(err)
	}
	timestamp, err := utils.HexToDecimal(hexTimestamp)
	if err != nil {
		logger.Log.Warnln(err)
	}
	date, err := utils.UnixToTime(timestamp, "UTC")
	if err != nil {
		logger.Log.Warnln(err)
	}

	c.JSON(http.StatusOK, ResponseBlockTimestamp{Timestamp: timestamp, HexTimestamp: hexTimestamp,
		Date: date})
}
