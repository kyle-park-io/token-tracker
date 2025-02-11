package main

import (
	"os"
	"time"

	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {

	// env
	env := "dev"
	viper.Set("ENV", env)
	switch env {
	case "dev":
		os.Setenv("ROOT_PATH", "/home/kyle/code/token-tracker/src")
		viper.Set("ROOT_PATH", "/home/kyle/code/token-tracker/src")

		os.Setenv("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
		viper.Set("LOG_PATH", viper.GetString("ROOT_PATH")+"/logs")
	case "prod":
		os.Setenv("ROOT_PATH", "/app")
		viper.Set("ROOT_PATH", "/app")

		os.Setenv("LOG_PATH", viper.GetString("ROOT_PATH")+"/../data/logs")
		viper.Set("LOG_PATH", viper.GetString("ROOT_PATH")+"/../data/logs")
	}

	// zap logger
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	// config
	os.Setenv("CONFIG_PATH", viper.GetString("ROOT_PATH")+"/configs/config.yaml")
	if err := config.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
	if env == "prod" {
		time.Sleep(10 * time.Second)
	}

	// Create the root command
	var rootCmd = &cobra.Command{
		Use:   "blockchain-tool",
		Short: "A blockchain utility tool with multiple server functionalities",
		Long:  `This application provides blockchain-related server functionalities, including block timestamp collection and transfer tracking, using the Cobra CLI library.`,
	}

	// Create the "block-timestamp" command
	var blockTimestampCmd = &cobra.Command{
		Use:   "block-timestamp",
		Short: "Start the server for collecting block timestamps",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartBlockTimestampServer()
		},
	}

	// Create the "transfer-tracker" command
	var transferTrackerCmd = &cobra.Command{
		Use:   "transfer-tracker",
		Short: "Start the server for tracking token transfers",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartTransferTrackerServer()
		},
	}

	// Create the "dev" command
	var devCmd = &cobra.Command{
		Use:   "dev",
		Short: "Start the server for executing experimental functions",
		Run: func(cmd *cobra.Command, args []string) {
			server.StartDevServer()
		},
	}

	// Add commands to the root command
	rootCmd.AddCommand(blockTimestampCmd)
	rootCmd.AddCommand(transferTrackerCmd)
	rootCmd.AddCommand(devCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Error(err)
	}
}
