package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kyle-park-io/token-tracker/integrated"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/tracker"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/gin-gonic/gin"
)

type ResponseBlockPosition struct {
	BlockTimestamp ResponseBlockTimestamp `json:"blockTimestamp"`
	BlockPosition  tracker.BlockPosition  `json:"blockPosition"`
}

// GetBlockPosition godoc
// @Summary Retrieve block position by date
// @Description Fetches the block position for a specific date in the given timezone. Returns block timestamp and position details.
// @Tags Block
// @Produce json
// @Param date query string true "Date in 'YYYY-MM-DD' format"
// @Success 200 {object} ResponseBlockPosition
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /get/getBlockPosition [get]
func GetBlockPosition(c *gin.Context) {

	date := c.Query("date")
	// Check if the input matches the format "YYYY-MM-DD"
	const layout = "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'date' must be in 'YYYY-MM-DD' format (e.g., '2024-01-01')",
		})
		return
	}

	// input := c.Query("timestamp")
	// timestamp := int64(0)
	// hexTimestamp := ""
	// // Check if the input is a decimal number
	// if n, err := strconv.Atoi(input); err == nil {
	// 	timestamp = int64(n)
	// 	hexTimestamp = utils.DecimalToHex(int64(n))
	// }
	// // Check if the input is a hexadecimal number
	// if strings.HasPrefix(input, "0x") {
	// 	bigInt := new(big.Int)
	// 	if _, success := bigInt.SetString(input[2:], 16); success {
	// 		timestamp, _ = utils.HexToDecimal(input)
	// 		hexTimestamp = input
	// 	} else {
	// 		// Return an error if the input is not a valid hexadecimal number
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "'number' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
	// 		})
	// 		return
	// 	}
	// }

	position, err := integrated.GetBlockPosition(date, "UTC")
	if err != nil {
		logger.Log.Warnln(err)
	}
	t, _ := utils.ConvertYearMonthDayToTimezone(date, "UTC")
	timestamp, _ := utils.TimeToUnix(t, "UTC")
	hexTimestamp := utils.DecimalToHex(timestamp)

	response := ResponseBlockPosition{
		BlockTimestamp: ResponseBlockTimestamp{Timestamp: timestamp, HexTimestamp: hexTimestamp, Date: t},
		BlockPosition:  position}
	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	c.JSON(http.StatusOK, response)
}
