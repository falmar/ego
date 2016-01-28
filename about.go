// About Controller
// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// About Page
func About(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	tpl := template.New("about.gohtml")

	tpl = tpl.Funcs(template.FuncMap{
		"MyFunc": func() string {
			return "I am a funny Func!"
		},
	})

	tpl, err = tpl.ParseFiles("templates/about.gohtml", "templates/inner.gohtml", "templates/menu.gohtml")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	UserID, err := strconv.ParseInt(getSession(w, r).Get("user-id"), 10, 64)
	User := getUser(UserID)

	err = tpl.Execute(w, BasicContext{
		Title: "About",
		Text:  "Hello World from About! a Mysterious function will be executed below",
		User:  User,
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// About Models
