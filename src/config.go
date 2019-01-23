package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	SiteName string
	Port int
}

func setupConfig() {
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Fatal(err)
	}
}
