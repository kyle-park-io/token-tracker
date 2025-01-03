package tracker

import (
	"encoding/json"
	"os"
)

func ReadBlockTimestampJSONFile(filePath string) ([]BlockTimestamp, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b []BlockTimestamp
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&b); err != nil {
		return nil, err
	}
	return b, nil
}

func MapTimestampByNumber(blockTimestamps []BlockTimestamp) map[string]Timestamp {
	timestampMap := make(map[string]Timestamp)
	for _, blockTimestamp := range blockTimestamps {
		timestampMap[blockTimestamp.Number] = blockTimestamp.Timestamp
	}
	return timestampMap
}

func MapHexTimestampByNumber(blockTimestamps []BlockTimestamp) map[string]string {
	timestampMap := make(map[string]string)
	for _, blockTimestamp := range blockTimestamps {
		timestampMap[blockTimestamp.Number] = blockTimestamp.Timestamp.Hex
	}
	return timestampMap
}

func MapStructByNumber(blockTimestamps []BlockTimestamp) map[string]struct{} {
	structMap := make(map[string]struct{})
	for _, blockTimestamp := range blockTimestamps {
		structMap[blockTimestamp.Number] = struct{}{}
	}
	return structMap
}
