package utils

import (
	"fmt"
	"math/big"
	"strings"
)

func DecimalToHex(bn int64) string {

	num := new(big.Int)
	num = num.SetInt64(bn)

	return fmt.Sprintf("0x%s", num.Text(16))
}

func HexToDecimal(bn string) (int64, error) {

	num := new(big.Int)
	num, success := num.SetString(bn[2:], 16)
	if !success {
		return 0, fmt.Errorf("failed to parse the hex string: %s", bn)
	}

	return num.Int64(), nil
}

func HexToDecimalString(bn string) (string, error) {

	num := new(big.Int)
	num, success := num.SetString(bn[2:], 16)
	if !success {
		return "", fmt.Errorf("failed to parse the hex string: %s", bn)
	}

	return num.String(), nil
}

func TrimLeadingZerosWithPrefix(hex string) string {
	// Ensure the input starts with "0x"
	if !strings.HasPrefix(hex, "0x") {
		return hex // Return as-is if it's not a valid hex string
	}
	// Remove "0x" prefix for processing
	trimmedHex := strings.TrimLeft(hex[2:], "0")
	// If all digits are removed (e.g., input was all zeros), return "0x0"
	if trimmedHex == "" {
		return "0x0"
	}
	// Re-add "0x" prefix
	return "0x" + trimmedHex
}
