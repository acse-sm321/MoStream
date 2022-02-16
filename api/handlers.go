package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// Handler create a new user
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

// Handler for user login
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("user_name")
	io.WriteString(w, username)
}
