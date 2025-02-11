package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kyle-park-io/token-tracker/logger"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetInt("database.port"),
		viper.GetString("database.sslmode"),
		viper.GetString("database.timezone"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	logger.Log.Infoln("✅ Successfully connected to PostgreSQL!")
	return nil
}

func CheckDBConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Log.Fatalf("Error getting database instance: %v", err)
	}

	// Ping check
	err = sqlDB.Ping()
	if err != nil {
		logger.Log.Infoln("❌ Database connection failed.")
	} else {
		logger.Log.Infoln("✅ Database connection is active.")
	}
}

func CreateDatabaseIfNotExists() {
	dbname := viper.GetString("database.dbname")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		"postgres",
		viper.GetInt("database.port"),
		viper.GetString("database.sslmode"),
		viper.GetString("database.timezone"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		logger.Log.Fatal(err)
	}

	if !exists {
		_, err = db.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Fatal(err)
		}
		logger.Log.Infoln("Database created successfully")
	} else {
		logger.Log.Infoln("Database already exists")
	}
}
