package router

import (
	"encoding/json"
	"net/http"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/gin-gonic/gin"
)

// ResponseBlockNumber represents the block number response
type ResponseBlockNumber struct {
	BlockNumber    int64  `json:"blockNumber" example:"12345678"`    // Decimal block number
	HexBlockNumber string `json:"hexBlockNumber" example:"0xabcdef"` // Hexadecimal block number
}

// GetLatestBlockNumber godoc
// @Summary Retrieve the latest block number
// @Description Fetches the latest block number from the blockchain in both decimal and hexadecimal formats.
// @Tags Block
// @Produce json
// @Success 200 {object} ResponseBlockNumber
// @Failure 500 {object} ErrorResponse
// @Router /get/getLatestBlockNumber [get]
func GetLatestBlockNumber(c *gin.Context) {

	hexBlokNumber, err := get.GetBlockNumber()
	if err != nil {
		logger.Log.Warnln(err)
	}

	blockNumber, err := utils.HexToDecimal(string(hexBlokNumber))
	if err != nil {
		logger.Log.Warnln(err)
	}

	response := ResponseBlockNumber{BlockNumber: blockNumber, HexBlockNumber: string(hexBlokNumber)}
	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	c.JSON(http.StatusOK, response)
}
