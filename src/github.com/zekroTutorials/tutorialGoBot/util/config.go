package util

import (
	"os"
	"encoding/json"
)

type Config struct {
	Token string `json:"token"`
	Prefix string `json:"prefix"`
}

var config Config

func LoadConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)
	decoder.Decode(&config)
}

func GetConfig() *Config {
	return &config
}