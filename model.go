package main

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/niconosenzo/basic-blog/config"
)

var dbconfig *config.Config

func init() {
	dbconfig = config.GetConfig()
}

type User struct {
	ID       int
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type Article struct {
	ID    int
	Title string
	Body  string
}

//list all the articles
func listArticles() (articles []Article, err error) {
	db := dbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		article := Article{}
		if err = rows.Scan(&article.ID, &article.Title, &article.Body); err != nil {
			return
		}
		articles = append(articles, article)
	}
	rows.Close()
	return
}

//list all the users
func listUsers() (users []User, err error) {
	db := dbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.UserName, &user.Password, &user.First, &user.Last, &user.Role); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

//update article
func (article Article) updateArticle() {
	db := dbConn()
	defer db.Close()
	insForm, err := db.Prepare("UPDATE article SET title=?, body=? WHERE id=?")
	check(err)
	_, err = insForm.Exec(article.Title, article.Body, article.ID)
	check(err)
	insForm.Close()
	return
}

//look for article
func getArticleByTitle(id string) (article Article) {
	ID, err := strconv.Atoi(id)
	check(err)
	article = Article{}
	db := dbConn()
	defer db.Close()
	err = db.QueryRow("SELECT * FROM article WHERE id = ?", ID).Scan(&article.ID, &article.Title, &article.Body)
	check(err)
	return
}

//check if username already exists, true if exists:
func checkUserName(un string) bool {
	var count int
	db := dbConn()
	defer db.Close()
	row := db.QueryRow("SELECT count(*) FROM user WHERE username = ?", un)
	err := row.Scan(&count)
	check(err)

	if count > 0 {
		return true
	}
	return false
}

//Get user by username
func getUserByUN(un string) (user User) {
	user = User{}
	db := dbConn()
	defer db.Close()
	err := db.QueryRow("SELECT * FROM user WHERE username = ?", un).Scan(&user.ID, &user.UserName, &user.Password, &user.First, &user.Last, &user.Role)
	check(err)
	return
}

//create a new user
func (user *User) createDBUser() {
	db := dbConn()
	defer db.Close()
	statement := "insert into user ( id, username, password, first, last, role) values (NULL, ?, ?, ?, ?, ?)"
	insform, err := db.Prepare(statement)
	check(err)
	_, err = insform.Exec(user.UserName, user.Password, user.First, user.Last, user.Role)
	check(err)
	insform.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	return
}
