package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"token-tracker/logger"
)

// SaveJSONToFile saves JSON data to a file with indentation
func SaveJSONToFile(data interface{}, filename string) error {
	// Convert the data to JSON with indentation
	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	if _, err := file.Write(formattedJSON); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	logger.Log.Infof("JSON data has been written to %s\n", filename)
	return nil
}
