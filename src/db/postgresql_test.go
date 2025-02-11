package db

import (
	"testing"

	"github.com/kyle-park-io/token-tracker/internal/config"
)

// go test -v -run TestInit
func TestInit(t *testing.T) {

	config.SetDevEnv()

	InitDB()

	t.Log("🔍 Checking database connection...")
	CheckDBConnection()
}
