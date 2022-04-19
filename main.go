package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("page/*.html"))
}

func calc(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "calc.html", nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}

	hp := http.FileServer(http.Dir("style"))

	http.Handle("/style/", http.StripPrefix("/style", hp))
	http.HandleFunc("/", calc)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("Cannot connect server to " + port)
	}
}
