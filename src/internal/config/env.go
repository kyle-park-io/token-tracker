package config

import (
	"os"

	"github.com/kyle-park-io/token-tracker/logger"

	"github.com/spf13/viper"
)

func SetDevEnv() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("ROOT_PATH", "/home/kyle/code/token-tracker/src")
	viper.Set("ROOT_PATH", "/home/kyle/code/token-tracker/src")

	os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
	if err := InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}