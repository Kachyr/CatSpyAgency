package main

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// .environment file variables
const (
	logLevelEnv = "LOG_LEVEL"
	// local db
	dbReadURLEnv  = "DB_READ_URL"
	dbWriteURLEnv = "DB_WRITE_URL"
	//
	ginPortEnv = "GIN_PORT"
)

type config struct {
	LogLevel zerolog.Level
	Database *dbConfig
	GinPort  string
}

type dbConfig struct {
	ReadURL  string
	WriteURL string
}

func loadConfig() (*config, error) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// .env file does not exist
		viper.AutomaticEnv()
	} else {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			return nil, errors.Wrap(err, "error reading config")
		}
	}

	logLevel, err := zerolog.ParseLevel(strings.ToLower(viper.GetString(logLevelEnv)))
	if err != nil {
		return nil, errors.Wrap(err, "error parsing log level")
	}

	return &config{
		LogLevel: logLevel,
		Database: &dbConfig{
			ReadURL:  viper.GetString(dbReadURLEnv),
			WriteURL: viper.GetString(dbWriteURLEnv),
		},
		GinPort: ":" + viper.GetString(ginPortEnv),
	}, nil

}
