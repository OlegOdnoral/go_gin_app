package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	App string
	DB  database `toml:"database"`
}

type database struct {
	Addr     string
	User     string
	Password string
	Database string
}

func GetConfig() {
	var data config
	if _, err := toml.DecodeFile("src/config/config.toml", &data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", data)

}
