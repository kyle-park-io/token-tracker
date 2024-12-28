package get

import (
	"encoding/hex"
	"math/big"

	"github.com/kyle-park-io/token-tracker/utils"
)

type TransferEvent struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

func GetRandomBlockNumber() (string, error) {

	blockNumber, err := GetBlockNumber()
	if err != nil {
		return "", err
	}

	r, err := utils.RandomHexBelow(string(blockNumber))
	if err != nil {
		return "", err
	}

	return r, nil
}

func DecodeTransferLog(eventHash string, topics []string, data string) TransferEvent {

	// decode: indexed params
	from := "0x" + topics[1][26:] // Extract the last 40 characters (20 bytes of the address)
	to := "0x" + topics[2][26:]

	// decode: non-indexed params (uint256 value)
	dataBytes, _ := hex.DecodeString(data[2:])
	value := new(big.Int).SetBytes(dataBytes)

	return TransferEvent{From: from, To: to, Value: "0x" + value.Text(16)}
}
