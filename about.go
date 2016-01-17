// About Controller
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

// About Page
func About(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tpl := template.New("about.gohtml")

	tpl = tpl.Funcs(template.FuncMap{
		"MyFunc": func() string {
			return "I am a funny Func!"
		},
	})

	tpl, err := tpl.ParseFiles("templates/about.gohtml")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = tpl.Execute(w, BasicContext{
		Title: "About",
		Text:  "Hello World from About! a Mysterios function will be executed below",
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// About Models