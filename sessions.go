package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const sessionLength int = 300000

type session struct {
	un           string
	lastActivity time.Time
}

func getUser(w http.ResponseWriter, r *http.Request) User {
	// get cookie
	c, err := r.Cookie("blogcookie")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "blogcookie",
			Value: sID.String(),
		}

	}
	//c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u User
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = getUserByUN(s.un)
	}
	return u
}

//returns false if a user is not logged in
func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("access_token")
	if err != nil {
		return false
	}
	return true
	// s, ok := dbSessions[c.Value]
	// fmt.Println(s)
	// // if ok {
	// // 	s.lastActivity = time.Now()
	// // 	dbSessions[c.Value] = s
	// // }
	// _, ok = dbUsers[s.un]
	// fmt.Println(dbUsers[s.un])
	// // refresh session
	// c.MaxAge = sessionLength
	// http.SetCookie(w, c)
	// return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30000) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
