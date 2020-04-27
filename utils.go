package main

import (
	"log"
	"net/http"
)

// Render template tmplName
func renderer(w http.ResponseWriter, tmplName string, p interface{}) {
	err := tmpl.ExecuteTemplate(w, tmplName, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
