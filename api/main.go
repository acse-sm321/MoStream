package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddleHandler(r *httprouter.Router) http.Handler {
	m := middlewareHandler{}
	m.r = r
	return m
}

// hack the http handler and make our own handler
func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// for each api,check the session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

// RegisterHandlers init a router and return
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleHandler(r)
	http.ListenAndServe(":8080", mh)
}

// listen->RegisterHandlers->handlers
// Create thousands of goroutines on multiple cores

// handler->validate{requests,user}-> logical operations -> response. [API services layer]
// validation-> data structure and err handling
// session? what is session and how to use
// middleware defs(message,err), handlers dbops and response
