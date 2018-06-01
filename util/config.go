package util

import (
	"os"
	"encoding/json"
)

// Config-Struktur mit JSON Key Namen
type Config struct {
	Token string `json:"token"`
	Prefix string `json:"prefix"`
}

// Lokale Variable der Konfig, welche von den
// internen Packages über GetConfig() zugegriffen
// werden kann.
var config Config

// LoadConfig läd die datei "config.json", der JSON-Decoder
// interpretiert sie und speichert die Daten in einer Instanz
// der Config structure.
// Zudem wird der Pointer der Instanz returnt.
func LoadConfig() *Config {
	f, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)
	decoder.Decode(&config)
	return &config
}

// GetConfig returnt den Pointer der Instanz der Config.
func GetConfig() *Config {
	return &config
}