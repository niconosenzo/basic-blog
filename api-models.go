package main

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func getArticle(id string) {
	ID, err := strconv.Atoi(id)
	check(err)
	article := Article{}
	db := dbConn()
	defer db.Close()
	err = db.QueryRow("SELECT * FROM article WHERE id = ?", ID).Scan(&article.ID, &article.Title, &article.Body)

	check(err)
	return
}
