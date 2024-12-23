package configs

import (
	"errors"
	"os"

	"token-tracker/logger"

	"github.com/spf13/viper"
)

func InitConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return errors.New("CONFIG_PATH environment variable is not set")
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("Error reading config file: %v", err)
		return err
	}

	return nil
}
