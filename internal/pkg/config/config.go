package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Apikey string   `json:"apikey"`
	Shares []string `json:"shares"`
}

func LoadFile(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
