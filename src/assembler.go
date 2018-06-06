package main

import (
	"github.com/gobuffalo/packr"
	"log"
	"os"
	"io/ioutil"
)

var (
	htmlFolder packr.Box
	cssFolder  packr.Box
	jsFolder   packr.Box
)

func init() {
	htmlFolder = packr.NewBox("./html")
	cssFolder = packr.NewBox("./css")
	jsFolder = packr.NewBox("./js")
}

/*
 * Unpacks all of the web files on to the source folder.
 */

func assemble() {

	dirs := []string{"./html/", "./css/", "./js/"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, 0775)
		} else {
			os.RemoveAll(dir)
			os.Remove(dir)
			os.Mkdir(dir, 0775)
		}
	}
	for _, str := range htmlFolder.List() {
		file, err := htmlFolder.MustBytes(str)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./html/"+str, file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Loaded file %s.", str)
	}
	for _, str := range cssFolder.List() {
		file, err := cssFolder.MustBytes(str)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./css/"+str, file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Loaded file %s.", str)
	}
	for _, str := range jsFolder.List() {
		file, err := jsFolder.MustBytes(str)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./js/"+str, file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Loaded file %s.", str)
	}
}
