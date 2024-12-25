package configs

import (
	"os"

	"token-tracker/logger"
)

func SetEnv() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
	if err := InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
