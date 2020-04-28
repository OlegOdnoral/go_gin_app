package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	DB DatabaseConfig `toml:"database"`
}

//DatabaseConfig config struct description
type DatabaseConfig struct {
	Addr            string
	User            string
	Password        string
	Database        string
	ApplicationName string
}

//GetDBConfig read config file an return database config struct
func GetDBConfig() (DatabaseConfig, error) {
	var configData config
	if _, err := toml.DecodeFile("src/config/config.toml", &configData); err != nil {
		fmt.Println(322)
		fmt.Println(err)
		return DatabaseConfig{}, err
	}
	return configData.DB, nil
}
