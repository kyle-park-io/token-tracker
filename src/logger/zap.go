package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func InitLogger() {

	fileName := viper.GetString("LOG_PATH") + "/logs.log"
	// Configure lumberjack logger
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName, // Log file name
		MaxSize:    10,       // Maximum file size (in MB)
		MaxBackups: 5,        // Maximum number of backup files
		MaxAge:     30,       // Maximum retention days
		Compress:   true,     // Whether to compress old log files
	})

	// Configure zapcore for file
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel)

	// Configure zapcore for console
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.AddSync(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel)

	// Combine cores
	core := zapcore.NewTee(fileCore, consoleCore)

	// Create logger
	logger := zap.New(core)
	// logger, err := zap.NewDevelopment()
	// if err != nil {
	// 	panic(err)
	// }
	defer logger.Sync() // Ensure any buffered logs are flushed
	Log = logger.Sugar()
}
