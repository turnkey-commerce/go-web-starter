package config

import (
	"log"
	"os"
	"path"

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
	// Check if it exists at the top level and if not look in the config directory
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		configFile = path.Join("config", configFile)
	}
	if _, err := toml.DecodeFile(configFile, &Settings); err != nil {
		log.Fatal(err)
	}
}
