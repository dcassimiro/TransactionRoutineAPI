package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func Url() string {
	config := LoadConfiguration("config.json")
	// url := config.Database.Host + ":" + config.Database.Password + "@tcp(local-mysql:3306)/"
	url := config.Database.Host + ":" + config.Database.Password + "@/"
	return url
}
