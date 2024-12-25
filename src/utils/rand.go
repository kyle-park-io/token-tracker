package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// RandomHexBelow generates a random value below the given hexadecimal string.
func RandomHexBelow(hexStr string) (string, error) {
	// Remove "0x" prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Convert the hex string to a big.Int
	upperLimit, ok := new(big.Int).SetString(hexStr, 16)
	if !ok {
		return "", fmt.Errorf("invalid hexadecimal string: %s", hexStr)
	}

	// Generate a random value below the upper limit
	randomValue, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		return "", fmt.Errorf("failed to generate random value: %v", err)
	}

	// Convert the random value back to a hex string
	return fmt.Sprintf("0x%s", randomValue.Text(16)), nil
}
