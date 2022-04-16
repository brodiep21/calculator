package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("image/*.html"))
}

func calc(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "calc.html", nil)
}

// func calculator(w http.ResponseWriter, r *http.Request) {
// 	switch r.FormValue() {
// 	case "clear" :
// 	}
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}

	http.HandleFunc("/", calc)
	http.Handle("/image/", http.FileServer(http.Dir("/image/styles.css")))
	http.ListenAndServe(":"+port, nil)
}
