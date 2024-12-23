package main

import (
	"token-tracker/configs"
	"token-tracker/logger"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	if err := configs.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
