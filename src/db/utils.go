package db

import (
	"encoding/json"
	"os"
)

func ReadInterfaceJSONFile(filePath string) ([]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
