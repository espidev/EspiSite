package main

/*
    EspiSite - a quick and dirty CMS
    Copyright (C) 2019 EspiDev

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	SiteName string
	Port int
	AdminRoute string
	Secret string
	Domain string
}

const (
	ConfigLocation = "./config.toml"
	DefaultConfig  = `
SiteName = "espidev"
Port = 3000
AdminRoute = "/admin/"
Secret = "hithisisnice"
Domain = "localhost"
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
