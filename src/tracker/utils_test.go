package tracker

import "testing"

// go test -v -run TestReadBlockTimestampJSONFile
func TestReadBlockTimestampJSONFile(t *testing.T) {
	filePath := "/home/kyle/code/token-tracker/src/json/blockTimestamp/blockTimestamp.json"
	b, err := ReadBlockTimestampJSONFile(filePath)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Timestamps:\n %+v", b)
}

// go test -v -run TestMapTimestampByNumber
func TestMapTimestampByNumber(t *testing.T) {
	filePath := "/home/kyle/code/token-tracker/src/json/blockTimestamp/blockTimestamp.json"
	b, err := ReadBlockTimestampJSONFile(filePath)
	if err != nil {
		t.Error(err)
	}

	timestampMap := MapTimestampByNumber(b)
	t.Logf("Timestamps:\n %+v", timestampMap)
}
