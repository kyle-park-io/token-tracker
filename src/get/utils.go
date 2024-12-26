package get

import "token-tracker/utils"

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
