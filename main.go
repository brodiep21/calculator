package main

import (
	"log"
	"net/http"
	"os"
)

// var templ *template.Template

// func init() {
// 	templ = template.Must(template.ParseGlob("webpage/*.html"))
// }

// func calc(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "calc.html", nil)
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Setting default port to %s", port)
	}

	// http.HandleFunc("/", calc)
	http.Handle("/", http.FileServer(http.Dir("webpage/")))
	http.ListenAndServe(":"+port, nil)
}
