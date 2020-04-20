package main

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

var getRidtxt = template.FuncMap{"grtxt": removeTxt}

func removeTxt(s string) string {
	return strings.Trim(s, ".txt")
}

//Extract title and check its format against validPath regex
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

// Render template tmplName
func renderer(w http.ResponseWriter, tmplName string, p interface{}) {
	err := tmpl.ExecuteTemplate(w, tmplName, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//Wrap handlers to avoid repeating checks
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
