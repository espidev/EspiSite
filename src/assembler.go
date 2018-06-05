package main

import (
	"github.com/gobuffalo/packr"
	"log"
)

var (
	htmlFolder packr.Box
	cssFolder  packr.Box
	jsFolder   packr.Box
)

func init() {
	htmlFolder = packr.NewBox("./html")
	cssFolder = packr.NewBox("./css")
	jsFolder =  packr.NewBox("./js")
}

/*
 * Unpacks all of the web files on to the source folder.
 */

func assemble() {
	for _, str := range htmlFolder.List() {
		file, err := htmlFolder.MustBytes(str)
		if err != nil {
			log.Fatal(err)
		}
		
	}
}
