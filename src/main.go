package main

import (
	"os"

	"github.com/kyle-park-io/token-tracker/internal/config"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Hi! i'm token tracker.")

	env := "dev"
	switch env {
	case "dev":
		os.Setenv("ROOT_PATH", "/home/kyle/code/token-tracker/src")
		viper.Set("ROOT_PATH", "/home/kyle/code/token-tracker/src")

		os.Setenv("CONFIG_PATH", "/home/kyle/code/token-tracker/src/configs/config.yaml")
		if err := config.InitConfig(); err != nil {
			logger.Log.Fatalf("Check Errors, %v", err)
		}
	case "prod":
		os.Setenv("ROOT_PATH", "/app")
		viper.Set("ROOT_PATH", "/app")

		os.Setenv("CONFIG_PATH", "/app/configs/config.yaml")
		if err := config.InitConfig(); err != nil {
			logger.Log.Fatalf("Check Errors, %v", err)
		}
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
			// server.StartBlockTimestampServer()
			server.StartBlockTimestampServer2()
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

	// Add commands to the root command
	rootCmd.AddCommand(blockTimestampCmd)
	rootCmd.AddCommand(transferTrackerCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Error(err)
	}
}
