package config

import (
	"github.com/Brandon-lz/gmqtt/utils"
	"github.com/BurntSushi/toml"
)

var Config ConfigDF

// config.toml
// [service]
// port = 8080

// [logging]
// level = "info"

// [redis]
// host = "localhost"
// port = 6379
// db = 0
// password = ""




type ConfigDF struct {
	Service struct {
		Port int `json:"port"`
	} `json:"service"`
	Logging struct {
		Level string `json:"level"`
	} `json:"logging"`
}


func LoadConfig(tomlFilePath string) (error) {
	var tomlData map[string]interface{}
    if _, err := toml.DecodeFile(tomlFilePath, &tomlData); err != nil {
        return err
    }
    utils.DeserializeData(tomlData, &Config)
	return nil
}