// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Home)
	router.GET("/About", About)
	router.GET("/Contact", Contact)
	router.POST("/Contact", ContactPost)
	router.ServeFiles("/src/*filepath", http.Dir("./public"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// BasicContext for templates
type BasicContext struct {
	Title string
	Text  string
}
