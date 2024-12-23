package utils

import (
	"fmt"
	"math/big"
)

func HexToDecimal(bn string) (string, error) {

	num := new(big.Int)
	num, success := num.SetString(bn[2:], 16)
	if !success {
		return "", fmt.Errorf("failed to parse the hex string: %s", bn)
	}

	return num.String(), nil
}
