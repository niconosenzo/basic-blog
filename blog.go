package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	tmpl              *template.Template
	dbUsers           = map[string]User{}    // user ID, user
	dbSessions        = map[string]session{} // session ID, session
	dbSessionsCleaned time.Time
)

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

//save and load

func main() {

	//http.HandleFunc("/view/", viewHandler)
	//cleanSessions()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/edit/{id}", editHandler)
	rtr.HandleFunc("/", listHandler)
	rtr.HandleFunc("/list/", listHandler)
	rtr.HandleFunc("/signup/", signupHandler)
	rtr.HandleFunc("/login/", loginHandler)
	rtr.HandleFunc("/api/", apiHandler)
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
