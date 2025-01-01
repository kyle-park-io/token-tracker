package router

import (
	"encoding/json"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/kyle-park-io/token-tracker/get"
	"github.com/kyle-park-io/token-tracker/get/call"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"
	"github.com/kyle-park-io/token-tracker/wss"

	"github.com/gin-gonic/gin"
)

type ResponseBalance struct {
	Balance    string `json:"balance" example:"100000000"`
	HexBalance string `json:"hexBalance" example:"0x5f5e100"`
}

// GetETHBalance godoc
// @Summary Retrieve the balance of an Ethereum account
// @Description Fetches the balance of a given Ethereum account, returning it in both decimal and hexadecimal formats.
// @Tags Balance
// @Produce json
// @Param account query string true "Ethereum account address"
// @Param tag query string false "(optional) Block tag (e.g., 'latest', 'earliest', or a block number in decimal/hexadecimal)"
// @Success 200 {object} ResponseBalance "Successfully retrieved the Ethereum balance"
// @Failure 400 {object} ErrorResponse "Invalid input, such as incorrect Ethereum address or tag"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /get/getETHBalance [get]
func GetETHBalance(c *gin.Context) {

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

	tag := c.Query("tag")
	tag2 := ""
	switch tag {
	case "":
		tag2 = "latest"
	case "latest":
		tag2 = "latest"
	case "earliest":
		tag2 = "earliest"
	default:
		// Check if the input is a decimal number
		if n, err := strconv.Atoi(tag); err == nil {
			// Convert the decimal number to a hexadecimal string
			tag2 = utils.DecimalToHex(int64(n))
			break
		}

		// Check if the input is a valid hexadecimal number
		if strings.HasPrefix(tag, "0x") {
			bigInt := new(big.Int)
			if _, success := bigInt.SetString(tag[2:], 16); success {
				// If valid, use the hexadecimal input as is
				tag2 = tag
			} else {
				// Respond with an error if the input is invalid
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "'tag' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
				})
				return
			}
		} else {
			// Respond with an error if the input is invalid
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'tag' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
			})
			return
		}
	}

	b, err := get.GetBalance(account, tag2)
	if err != nil {
		logger.Log.Warnln(err)
	}
	balanceHex := utils.TrimLeadingZerosWithPrefix(string(b))
	balance, _ := utils.HexToDecimalString(balanceHex)

	response := ResponseBalance{Balance: balance, HexBalance: balanceHex}
	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	// Send the block data as a JSON response
	c.JSON(http.StatusOK, response)
}

// GetERC20Balance godoc
// @Summary Retrieve the balance of an ERC-20 token for a given Ethereum account
// @Description Fetches the balance of a specific ERC-20 token for a given Ethereum account, returning it in both decimal and hexadecimal formats.
// @Tags Balance
// @Produce json
// @Param account query string true "Ethereum account address"
// @Param tokenAddress query string true "ERC-20 token contract address"
// @Param tag query string false "(optional) Block tag (e.g., 'latest', 'earliest', or a block number in decimal/hexadecimal)"
// @Success 200 {object} ResponseBalance "Successfully retrieved the ERC-20 token balance"
// @Failure 400 {object} ErrorResponse "Invalid input, such as incorrect Ethereum address, token address, or tag"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /get/getERC20Balance [get]
func GetERC20Balance(c *gin.Context) {

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

	tokenAddress := c.Query("tokenAddress")
	matched, _ = regexp.MatchString(pattern, tokenAddress)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'tokenAddress' must be a valid Ethereum address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
		})
		return
	}

	tag := c.Query("tag")
	tag2 := ""
	switch tag {
	case "":
		tag2 = "latest"
	case "latest":
		tag2 = "latest"
	case "earliest":
		tag2 = "earliest"
	default:
		// Check if the input is a decimal number
		if n, err := strconv.Atoi(tag); err == nil {
			// Convert the decimal number to a hexadecimal string
			tag2 = utils.DecimalToHex(int64(n))
			break
		}

		// Check if the input is a valid hexadecimal number
		if strings.HasPrefix(tag, "0x") {
			bigInt := new(big.Int)
			if _, success := bigInt.SetString(tag[2:], 16); success {
				// If valid, use the hexadecimal input as is
				tag2 = tag
			} else {
				// Respond with an error if the input is invalid
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "'tag' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
				})
				return
			}
		} else {
			// Respond with an error if the input is invalid
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "'tag' must be a valid decimal or hexadecimal string (e.g., '123' or '0x1A')",
			})
			return
		}
	}

	// Example: ERC-20 balanceOf(address)
	methodName := "balanceOf"
	paramTypes := []string{"address"}
	params := []interface{}{account}
	callData, _ := call.CreateCallData(methodName, paramTypes, params)

	txArgs := map[string]interface{}{"to": tokenAddress, "data": callData}
	b, err := get.GetCallBalance(txArgs, tag2)
	if err != nil {
		logger.Log.Warnln(err)
	}
	balanceHex := utils.TrimLeadingZerosWithPrefix(string(b))
	balance, _ := utils.HexToDecimalString(balanceHex)

	response := ResponseBalance{Balance: balance, HexBalance: balanceHex}

	jsonData, _ := json.Marshal(response)
	// ws.GlobalLogChannel <- string(jsonData)
	wss.GlobalLogChannel <- string(jsonData)

	// Send the block data as a JSON response
	c.JSON(http.StatusOK, response)
}
