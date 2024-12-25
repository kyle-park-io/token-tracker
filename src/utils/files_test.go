package utils

import (
	"testing"

	"token-tracker/configs"
)

// go test -v -run TestFileExists
func TestFileExists(t *testing.T) {

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	exists, err := FileExists(filePath)
	if err != nil {
		t.Error(err)
	}

	if !exists {
		t.Log("File isn't exists!")
	} else {
		t.Log("File exists!")
	}
}

// go test -v -run TestEnsureFileExists
func TestEnsureFileExists(t *testing.T) {

	configs.SetEnv()

	filePath := "/home/kyle/code/token-tracker/src/get/json/block.json"
	err := EnsureFileExists(filePath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Successfully ensure file exists.")
}
