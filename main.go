package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templ *template.Template

func calc(w http.ResponseWriter, r *http.Request) {
	templ, _ = template.ParseFiles("./webpage/calc.html")
	templ.Execute(w, "calc.html")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/webpage/styles.css")
	})

	http.HandleFunc("/", calc)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("Cannot connect server to " + port)
	}
}
