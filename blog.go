package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
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
	//tmpl      = template.Must(template.ParseGlob("templates/*"))
	//tmpl      = template.Must(template.New("").Funcs(getRidtxt).ParseGlob("templates/*"))
	validPath = regexp.MustCompile("^/(edit|save|view|list)/([a-zA-Z0-9]+)$")
)

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

//save and load

func main() {
	// creating the page
	// p := &Page{Title: "TestPage", Body: []byte("This is the body of the test page.")}

	// //save the page and check for errors
	// if err := p.Save(); err != nil {
	// 	log.Fatal(err)
	// }

	//http.HandleFunc("/view/", viewHandler)
	//cleanSessions()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/edit/{id}", editHandler)
	rtr.HandleFunc("/", listHandler)
	rtr.HandleFunc("/list/", listHandler)
	rtr.HandleFunc("/signup/", signupHandler)
	rtr.HandleFunc("/login/", loginHandler)
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
