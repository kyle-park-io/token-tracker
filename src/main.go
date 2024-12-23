package main

import (
	"os"

	"token-tracker/configs"
	"token-tracker/logger"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
	if err := configs.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
