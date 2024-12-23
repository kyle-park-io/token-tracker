package configs

import (
	"token-tracker/logger"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("Error reading config file: %v", err)
		return err
	}

	return nil
}
