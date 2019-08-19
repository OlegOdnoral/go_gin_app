package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	DB DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Addr            string
	User            string
	Password        string
	Database        string
	ApplicationName string
}

func GetDBConfig() (DatabaseConfig, error) {
	var configData config
	if _, err := toml.DecodeFile("src/config/config.toml", &configData); err != nil {
		fmt.Println(err)
		return DatabaseConfig{}, err
	}
	return configData.DB, nil
}
