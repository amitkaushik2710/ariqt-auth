package config

import (
	"ariqt-auth/pkg/logger"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Port int `json:"port"`
}

func GetConfig() *Config {
	var config Config

	configFile, err := os.Open("config/config.json")
	if err != nil {
		logger.Logger.Fatal("GetConfig: failed to open config json", zap.Error(err))
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config
}
