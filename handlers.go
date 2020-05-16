package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//Handlers

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// if alreadyLoggedIn(w, r) {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }
	var u User
	// process form submission
	if r.Method == http.MethodPost {
		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		role := r.FormValue("role")
		// username taken?
		if checkUserName(un) {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		// sID, _ := uuid.NewV4()
		// c := &http.Cookie{
		// 	Name:  "blogcookie",
		// 	Value: sID.String(),
		// 	Path:  "/",
		// }
		//c.MaxAge = sessionLength
		// http.SetCookie(w, c)
		// dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		//setting ID to 0, it will be assigned by the DB engine
		u = User{0, un, bs, f, l, role}
		u.createDBUser()
		// redirect
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	showSessions()
	renderer(w, "Signup", u)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	showSessions()
	// if alreadyLoggedIn(w, r) {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }
	var u User
	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		// is there a username?
		ok := checkUserName(un)
		if !ok {
			http.Error(w, "Username doesn't exist", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		u = getUserByUN(un)
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "blogcookie",
			Value: sID.String(),
			Path:  "/",
		}
		//c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}
	showSessions()
	renderer(w, "Login", u)
}

func editHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	varID := vars["id"]
	fmt.Println(varID)
	// if !alreadyLoggedIn(w, r) {
	// 	http.Redirect(w, r, "/signup", http.StatusSeeOther)
	// 	return
	// }
	p := getArticleByTitle(varID)

	if r.Method == http.MethodPost {
		p.Body = r.FormValue("body")
		p.Title = r.FormValue("title")
		p.updateArticle()
		http.Redirect(w, r, "/", http.StatusFound)
	}

	renderer(w, "Form", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	// p := getArticleByTitle(title)
	// p.Body = r.FormValue("body")
	// p.Title = r.FormValue("title")
	// p.Save()
	// http.Redirect(w, r, "/list", http.StatusFound)
	return

}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("blogcookie")
	// delete the session
	delete(dbSessions, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 180) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("within the list")
	showSessions()
	//u := getUser(w, r)
	if !alreadyLoggedIn(w, r) {
		fmt.Println("redirecting you")
		fmt.Println(r.Cookie("blogcookie"))
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	pages, err := listArticles()
	if err != nil {
		log.Fatal(err)
	}

	renderer(w, "List", pages)

}
