package utils

import (
	"testing"
)

// go test -v -run TestFindGoModPath
func TestFindGoModPath(t *testing.T) {

	root, err := FindGoModPath()
	if err != nil {
		t.Error("Error:", err)
		return
	}
	t.Log("Go Module Root Path:", root)
}
