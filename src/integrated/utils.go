package integrated

import "token-tracker/utils"

func splitBlockRange(toBlockNumber, fromBlockNumber string, maxResults int64) [][]string {

	toBlock, _ := utils.HexToDecimal(toBlockNumber)
	fromBlock, _ := utils.HexToDecimal(fromBlockNumber)

	var ranges [][]string
	for toBlock >= fromBlock {
		startBlock := toBlock - maxResults + 1
		if startBlock < fromBlock {
			startBlock = fromBlock
		}
		ranges = append(ranges, []string{utils.DecimalToHex(startBlock), utils.DecimalToHex(toBlock)})
		toBlock = startBlock - 1
	}
	return ranges
}
