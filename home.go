// Home Controller
// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Home Page
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/home.gohtml", "templates/menu.gohtml")
	if err != nil {
		return
	}

	UserID, err := strconv.ParseInt(getSession(w, r).Get("user-id"), 10, 64)
	User := getUser(UserID)

	err = tpl.Execute(w, BasicContext{
		Title: "Home",
		Text:  "Hello World",
		User:  User,
	})

	if err != nil {
		log.Println(err)
	}
}

// Home Models
