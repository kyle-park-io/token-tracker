package call

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/sha3"
)

// Keccak256 calculates the Keccak256 hash of the input data
func Keccak256ToByte(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// Keccak256 calculates the Keccak256 hash of the input data
func Keccak256ToString(data []byte) string {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// EncodeUint256 encodes a uint256 value into a 32-byte hex representation
func EncodeUint256(value *big.Int) string {
	return fmt.Sprintf("%064x", value)
}

// EncodeAddress encodes an Ethereum address into a 32-byte hex representation
func EncodeAddress(address string) (string, error) {
	// Remove the '0x' prefix if present
	address = strings.TrimPrefix(address, "0x")
	if len(address) != 40 {
		return "", fmt.Errorf("invalid address length: %s", address)
	}
	// Pad to 32 bytes
	return fmt.Sprintf("%064s", address), nil
}

// CreateCallData creates the call data for a smart contract function
func CreateCallData(methodName string, paramTypes []string, params []interface{}) (string, error) {
	// Construct the method signature
	signature := fmt.Sprintf("%s(%s)", methodName, strings.Join(paramTypes, ","))

	// Compute the Keccak256 hash of the method signature and take the first 4 bytes
	hash := Keccak256ToByte([]byte(signature))
	functionSelector := hex.EncodeToString(hash[:4])

	// Buffer to hold the encoded parameters
	var encodedParams bytes.Buffer

	// Encode each parameter
	for i, param := range params {
		paramType := paramTypes[i]
		switch paramType {
		case "uint256":
			value, ok := param.(*big.Int)
			if !ok {
				return "", fmt.Errorf("parameter %d must be *big.Int for type uint256", i)
			}
			encodedParams.WriteString(EncodeUint256(value))
		case "address":
			address, ok := param.(string)
			if !ok {
				return "", fmt.Errorf("parameter %d must be string for type address", i)
			}
			encodedAddress, err := EncodeAddress(address)
			if err != nil {
				return "", err
			}
			encodedParams.WriteString(encodedAddress)
		default:
			return "", fmt.Errorf("unsupported parameter type: %s", paramType)
		}
	}

	// Combine function selector and encoded parameters
	return "0x" + functionSelector + encodedParams.String(), nil
}
