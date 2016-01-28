package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func setLogHandlers(r *httprouter.Router) {
	r.GET("/Login", Login)
	r.POST("/Login", LoginP)
	r.GET("/Logout", Logout)
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/login.gohtml", "templates/menu.gohtml")
	if err != nil {
		log.Println(err)
		return
	}

	tpl.Execute(w, BasicContext{Title: "Login", Text: "Use your email to log in:", User: getUser(0)})
	if err != nil {
		log.Println(err)
		return
	}

}

func LoginP(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()

	Email := r.PostFormValue("Email")

	if Email == "" {
		w.Write([]byte("Email can not be empty string"))
		return
	}

	db, err := getConn()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := db.Prepare("SELECT ID FROM User WHERE Email = ?")
	if err != nil {
		log.Println(err)
		return
	}

	var UserID int64

	err = stmt.QueryRow(Email).Scan(&UserID)
	if err != nil || UserID <= 0 {
		w.Write([]byte("Invalid email"))
		return
	}

	Lifetime := time.Duration(15) * time.Minute

	getSession(w, r).Set("user-id", strconv.FormatInt(UserID, 10), time.Now().Add(Lifetime))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Lifetime := time.Duration(-1) * time.Minute
	getSession(w, r).Set("user-id", "", time.Now().Add(Lifetime))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
