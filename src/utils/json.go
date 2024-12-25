package utils

import (
	"bufio"
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

// SaveJSONDirect saves JSON data directly to a file without buffering.
// It creates a temporary file, writes the JSON data directly, and then renames
// the temporary file to the target filename. Suitable for smaller data writes.
func SaveJSONDirect(data interface{}, filename string) error {
	// Convert the data to JSON with indentation
	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "temp-json-*.json")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure temp file is removed

	// Write the JSON data to the temporary file
	if _, err := tempFile.Write(formattedJSON); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Ensure all data is written to disk
	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}
	tempFile.Close()

	// Replace the original file with the temporary file
	if err := os.Rename(tempFile.Name(), filename); err != nil {
		return fmt.Errorf("failed to replace original file: %w", err)
	}

	fmt.Printf("JSON data has been written to %s\n", filename)
	return nil
}

// SaveJSONBuffered saves JSON data to a file using a buffered writer.
// It creates a temporary file, writes the JSON data to a buffer first,
// flushes the buffer to the file, and then renames the temporary file
// to the target filename. Ideal for larger data writes to improve performance.
func SaveJSONBuffered(data interface{}, filename string) error {
	// Convert the data to JSON with indentation
	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "temp-json-*.json")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure temp file is removed

	// Create a buffered writer
	bufferedWriter := bufio.NewWriter(tempFile)

	// Write JSON data to the buffered writer
	if _, err := bufferedWriter.Write(formattedJSON); err != nil {
		return fmt.Errorf("failed to write to buffered writer: %w", err)
	}

	// Ensure all data in the buffer is written to disk
	if err := bufferedWriter.Flush(); err != nil {
		return fmt.Errorf("failed to flush buffered writer: %w", err)
	}

	// Ensure all data is synced to disk
	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}
	tempFile.Close()

	// Replace the original file with the temporary file
	if err := os.Rename(tempFile.Name(), filename); err != nil {
		return fmt.Errorf("failed to replace original file: %w", err)
	}

	fmt.Printf("JSON data has been written to %s\n", filename)
	return nil
}
