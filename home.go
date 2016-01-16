// Home Controller
// link       https://github.com/falmar/ego
// author     David Lavieri (falmar) <daviddlavier@gmail.com>
// copyright  2016 David Lavieri
// license    http://opensource.org/licenses/MIT The MIT License (MIT)

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home method or controller
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Home - Hello World!"))
}

// Home Models
