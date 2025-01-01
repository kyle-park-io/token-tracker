package config

import (
	"os"

	"github.com/kyle-park-io/token-tracker/logger"

	"github.com/spf13/viper"
)

func SetDevEnv() {
	os.Setenv("ROOT_PATH", "/home/kyle/code/token-tracker/src")
	viper.Set("ROOT_PATH", "/home/kyle/code/token-tracker/src")

	os.Setenv("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	viper.Set("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")

	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	os.Setenv("CONFIG_PATH", viper.GetString("ROOT_PATH")+"/configs/config.yaml")
	if err := InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
