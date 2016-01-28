package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func setMySQLHandlers(r *httprouter.Router) {
	r.GET("/MySQL", MySQL) // Read

	// GET Handlers
	r.GET("/Create", Create)
	r.GET("/Update/:ID", Update)
	r.GET("/Delete/:ID", Delete)

	//Post Handlers
	r.POST("/Create", CreateP)
	r.POST("/Update/:ID", UpdateP)
}

// MySQL Examples
func MySQL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/mysql/mysql.gohtml", "templates/menu.gohtml")
	if err != nil {
		log.Println(err)
		return
	}

	db, err := getConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM User LIMIT ?,?")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(0, 15)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	Users := make([]User, 0, 15)

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.RegisteredDate)
		if err != nil {
			log.Println(err)
			return
		}
		Users = append(Users, *user)
	}

	UserID, err := strconv.ParseInt(getSession(w, r).Get("user-id"), 10, 64)
	User := getUser(UserID)

	err = tpl.Execute(w, MySQLContext{
		Title: "MySQL Examples",
		Text:  "MySQL - CRUD",
		Users: Users,
		User:  User,
	})

	if err != nil {
		log.Println(err)
		return
	}
}

// Create a new Record
func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/mysql/create.gohtml", "templates/menu.gohtml")
	if err != nil {
		log.Println(err)
		return
	}

	UserID, err := strconv.ParseInt(getSession(w, r).Get("user-id"), 10, 64)
	User := getUser(UserID)

	tpl.Execute(w, BasicContext{Title: "Create MySQL Record", Text: "Insert a new record to User table", User: User})
}

// Update a Record
func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID, err := strconv.ParseInt(ps.ByName("ID"), 10, 64)

	if err != nil {
		http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
		return
	}

	tpl, err := template.ParseFiles("templates/mysql/update.gohtml", "templates/menu.gohtml")
	if err != nil {
		log.Println(err)
		return
	}

	eUser := getUser(ID)

	err = tpl.Execute(w, map[string]string{
		"Title": "Update row",
		"Text":  fmt.Sprintf("Update row # %v", ID),
		"ID":    strconv.Itoa(int(eUser.ID)),
		"Name":  eUser.Name,
		"Email": eUser.Email,
	})

	if err != nil {
		log.Println(err)
		return
	}
}

// Delete a Record
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID := ps.ByName("ID")

	if ID == "" {
		http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
		return
	}

	db, err := getConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM User WHERE ID = ?")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(ID)
	if err != nil {
		log.Println(err)
		return
	}

	if rws, err := res.RowsAffected(); rws <= 0 || err != nil {
		log.Println("No Rows were affected")
		return
	}

	http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
}

// CreateP "Post" a new Record
func CreateP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	r.ParseForm()

	var Name, Email string = r.PostFormValue("Name"), r.PostFormValue("Email")

	if Name == "" || Email == "" {
		w.Write([]byte("Name or Email can not be empty string"))
		return
	}

	db, err := getConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO User (Name,Email,RegisteredDate) VALUES (?,?,?)")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(Name, Email, getTimeStamp())

	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
}

// UpdateP "Post" a Record
func UpdateP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ID := ps.ByName("ID")

	if ID == "" {
		http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
		return
	}

	r.ParseForm()

	var Name, Email string = r.PostFormValue("Name"), r.PostFormValue("Email")

	if Name == "" || Email == "" {
		w.Write([]byte("Name or Email can not be empty string"))
		return
	}

	db, err := getConn()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE User SET Name = ?, Email = ? WHERE ID = ?")

	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(Name, Email, ID)

	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/MySQL", http.StatusMovedPermanently)
}

func getConn() (*sql.DB, error) {
	return sql.Open("mysql", "ego@tcp(127.0.0.1:3307)/ego")
}

func getTimeStamp() string {
	return time.Now().Format("2006-01-02 03:04:05")
}

// MySQLContext for read
type MySQLContext struct {
	Title, Text string
	Users       []User
	User        *User
}

// User struct
type User struct {
	ID             int64
	Name, Email    string
	RegisteredDate string
}

func getUser(UserID int64) *User {

	User := &User{
		Name: "Guest",
	}

	if UserID > 0 {

		db, err := getConn()
		if err != nil {
			log.Println(err)
		}
		defer db.Close()

		stmt, err := db.Prepare("SELECT Name,Email,RegisteredDate FROM User WHERE ID = ?")
		if err != nil {
			log.Println(err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(UserID).Scan(&User.Name, &User.Email, &User.RegisteredDate)
		if err != nil {
			log.Println(err)
		}

		User.ID = UserID
	}

	return User
}
