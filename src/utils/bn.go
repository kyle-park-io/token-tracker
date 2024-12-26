package utils

import (
	"fmt"
	"math/big"
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
