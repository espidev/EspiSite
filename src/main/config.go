package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	SiteName string
	Port int
}

const (
	ConfigLocation = "./config.toml"
	DefaultConfig  = `
	SiteName = "espidev"
	Port = 3000
`)

func setupConfig() {
	if _, err := os.Stat(ConfigLocation); os.IsNotExist(err) {
		err := ioutil.WriteFile(ConfigLocation, []byte(DefaultConfig), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := toml.DecodeFile(ConfigLocation, &config); err != nil {
		log.Fatal(err)
	}
}
