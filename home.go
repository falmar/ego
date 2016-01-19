// Home Controller
// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home Page
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/home.gohtml", "templates/menu.gohtml")
	if err != nil {
		return
	}
	tpl.Execute(w, BasicContext{Title: "Home", Text: "Hello World"})
}

// Home Models
