package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) (err error) {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//get object
	return respondJSON(w, 200, "payload")
}
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	//update object
	return respondJSON(w, 200, "payload")
}
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//create object
	return respondJSON(w, 200, "payload")
}
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	//delete object
	return respondJSON(w, 200, "payload")
}
