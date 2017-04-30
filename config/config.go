package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var configFile = "config.toml"

// Settings contains the settings the Website from the config.toml file.
var Settings struct {
	Website struct {
		HTTPPort    string `valid:"int,required"`
		CookieKey   string `valid:"ascii,required"`
		SecureHTTPS bool   `valid:"bool"`
	}
}

func init() {
	if _, err := toml.DecodeFile(configFile, &Settings); err != nil {
		log.Fatal(err)
	}
}
