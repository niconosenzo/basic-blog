package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	//tmpl      = template.Must(template.ParseGlob("templates/*"))
	tmpl      = template.Must(template.New("").Funcs(getRidtxt).ParseGlob("templates/*"))
	validPath = regexp.MustCompile("^/(edit|save|view|list)/([a-zA-Z0-9]+)$")
)

//save and load
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("articles/"+filename, p.Body, 0600)
}

func loadPage(fn string) (*Page, error) {
	filename := fn + ".txt"
	body, err := ioutil.ReadFile("articles/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: fn, Body: body}, nil
}

func main() {
	// creating the page
	p := &Page{Title: "TestPage", Body: []byte("This is the body of the test page.")}

	//save the page and check for errors
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", listHandler)    
	http.HandleFunc("/list/", listHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
