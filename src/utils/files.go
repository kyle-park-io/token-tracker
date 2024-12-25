package utils

import (
	"fmt"
	"os"

	"token-tracker/logger"
)

// FileExists checks if the file exists and returns a boolean.
func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	// If no error, the file exists
	if err == nil {
		return true, nil
	}
	// If the error is because the file doesn't exist
	if os.IsNotExist(err) {
		return false, err
	}
	// Other errors (e.g., permission issues)
	return false, err
}

// EnsureFileExists checks if a file exists, creates it if it doesn't, and does nothing if it exists.
func EnsureFileExists(filePath string) error {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists, nothing to do
		logger.Log.Info("File exists. Proceeding without changes.")
		return nil
	}

	if os.IsNotExist(err) {
		// File does not exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()
		logger.Log.Info("File did not exist. It has been created.")
		return nil
	}

	// Unexpected error
	return fmt.Errorf("error checking file: %w", err)
}
