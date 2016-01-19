package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ContactContext for Contact Template
type ContactContext struct {
	Title string
	Text  string
	Post  PostData
}

// PostData to store Post Data self explanatory (?)
type PostData struct {
	Name    string
	Email   string
	Message string
}

// Contact show the form for the user
func Contact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("templates/contact.gohtml", "templates/menu.gohtml")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = tpl.Execute(w, ContactContext{
		Title: "Contact",
		Text:  "Contact us!",
		Post:  *new(PostData),
	})

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

// ContactPost process the form sent by the user
func ContactPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tpl, err := template.ParseFiles("templates/contact.gohtml")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tpl.Execute(w, ContactContext{
		Title: "Contact Post",
		Text:  "This is your data:",
		Post: PostData{
			Name:    r.PostFormValue("Contact[Name]"),
			Email:   r.PostFormValue("Contact[Email]"),
			Message: r.PostFormValue("Contact[Message]"),
		},
	})
}
