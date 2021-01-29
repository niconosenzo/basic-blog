package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func dbConn() (db *sql.DB) {
	dbDriver := dbconfig.DB.Dialect
	dbUser := dbconfig.DB.Username
	dbPass := dbconfig.DB.Password
	dbName := dbconfig.DB.Name
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
