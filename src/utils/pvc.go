package utils

import (
	"io"
	"os"

	"github.com/kyle-park-io/token-tracker/logger"

	"github.com/spf13/viper"
)

func CheckPVCData() error {
	logger.Log.Infoln("Starting CheckPVCData")

	root := viper.GetString("ROOT_PATH")

	pvcPath := root + "/../data/blockTimestamp.json"
	jsonPath := root + "/json/blockTimestamp/blockTimestamp.json"
	b, _ := FileExists(pvcPath)
	if b {
		logger.Log.Infoln("Block timestamp is existed in PVC. copy data.")
		copyPVCDataToServer(pvcPath, jsonPath)
	}
	return nil
}

// copyPVCDataToServer copy blocktimestap data from the specified PVC-mounted path
func copyPVCDataToServer(filePath, dst string) error {
	// Open the source file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the content
	_, err = io.Copy(destFile, file)
	if err != nil {
		return err
	}

	// Flush and close destination file
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
