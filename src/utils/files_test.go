package utils

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
)

// go test -v -run TestFileExists
func TestFileExists(t *testing.T) {

	filePath := "/home/kyle/code/token-tracker/src/json/test/block/block.json"
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

	config.SetDevEnv()

	filePath := "/home/kyle/code/token-tracker/src/json/test/block/block.json"
	err := EnsureFileExists(filePath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Successfully ensure file exists.")
}
