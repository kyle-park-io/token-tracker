package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/utils"

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

	infuraURL := viper.GetString("infura.https_endpoint")
	if strings.Contains(infuraURL, "InputYour") {
		return fmt.Errorf("change the Infura API URL: %s", infuraURL)
	}

	// check pvc data
	env := viper.GetString("ENV")
	if env == "prod" {
		err := utils.CheckPVCData()
		if err != nil {
			logger.Log.Error(err)
			return err
		}
	}

	return nil
}
