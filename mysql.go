package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func setMySQLHandlers(r *httprouter.Router) {
	r.GET("/MySQL", MySQL) // Read
	r.GET("/Create", Create)
	r.GET("/Update/:ID", Update)
	r.GET("/Delete/:ID", Delete)
}

// MySQL Examples
func MySQL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/mysql/mysql.gohtml", "templates/menu.gohtml")
	if err != nil {
		return
	}

	db, err := sql.Open("mysql", "ego@tcp(127.0.0.1:3307)/ego")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM User LIMIT ?,?")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(0, 15)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	Users := make([]User, 0, 15)

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegisteredDate)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		Users = append(Users, *user)
	}

	err = tpl.Execute(w, MySQLContext{
		Title: "MySQL Examples",
		Text:  "MySQL - CRUD",
		Users: Users,
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// Create a new Record
func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Update a Record
func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Delete a Record
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// MySQLContext for read
type MySQLContext struct {
	Title, Text string
	Users       []User
}

// User struct
type User struct {
	ID             int
	Name, Email    string
	RegisteredDate string
}
