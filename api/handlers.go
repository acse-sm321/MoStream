package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// TODO Handler create a new user
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}
