// Home Controller
// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// HomeContext Struct goes to the home template
type HomeContext struct {
	Title string
	Text  string
}

// Home method or controller
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		return
	}
	tpl.Execute(w, HomeContext{Title: "Home", Text: "Hello World"})
}

// Home Models
