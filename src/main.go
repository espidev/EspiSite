package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"html/template"
	"log"
)

type Variables struct {
	Time string
}

func main() {
	log.Printf("Starting EspiSite...")
	log.Printf("Assembling files...")

	assemble() //gather files

	log.Printf("Initializing routers...")

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/home", homeHandler)
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/site", http.StripPrefix("/site", fs))

	log.Printf("Webserver starting on port 8080.")

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := Variables{}
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	vars := Variables{
		Time: now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("./html/home.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}
