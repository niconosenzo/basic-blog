package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Handlers
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		// newP := &Page{Title: title}
		// tmpl.ExecuteTemplate(w, "Form", newP)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderer(w, "View", p)

}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderer(w, "Form", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	return

}

//Handlers
func listHandler(w http.ResponseWriter, r *http.Request) {

	pages, err := ioutil.ReadDir("articles/")
	if err != nil {
		log.Fatal(err)
	}

	renderer(w, "List", pages)
	// for _, page := range pages {
	// 	fmt.Println(page.Name())
	// }
	// p, err := loadPage(title)
	// if err != nil {
	// 	// newP := &Page{Title: title}
	// 	// tmpl.ExecuteTemplate(w, "Form", newP)
	// 	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	// 	return
	// }
	// renderer(w, "View", p)

}
