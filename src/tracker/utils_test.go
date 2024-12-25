package tracker

import "testing"

// go test -v -run TestReadJSONFile
func TestReadJSONFile(t *testing.T) {
	filePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp.json"
	b, err := readJSONFile(filePath)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Timestamps:\n %+v", b)
}

// go test -v -run TestMapTimestampByNumber
func TestMapTimestampByNumber(t *testing.T) {
	filePath := "/home/kyle/code/token-tracker/src/get/json/blockTimestamp.json"
	b, err := readJSONFile(filePath)
	if err != nil {
		t.Error(err)
	}

	timestampMap := mapTimestampByNumber(b)
	t.Logf("Timestamps:\n %+v", timestampMap)
}
