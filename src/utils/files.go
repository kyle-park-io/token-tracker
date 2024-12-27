package utils

import (
	"fmt"
	"os"
	"path/filepath"

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

// CreateFolder creates a folder if it doesn't exist
func CreateFolder(folderPath string) error {
	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create the folder (including parent directories if needed)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %w", err)
		}
		fmt.Println("Folder created:", folderPath)
	}

	return nil
}

// CreateFolderAndFile creates a folder if it doesn't exist
// and then creates a file inside that folder.
func CreateFolderAndFile(folderPath, fileName string) error {
	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create the folder (including parent directories if needed)
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %w", err)
		}
		fmt.Println("Folder created:", folderPath)
	}

	// Create the file inside the folder
	filePath := filepath.Join(folderPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Ensure the file is closed
	fmt.Println("File created:", filePath)

	return nil
}
