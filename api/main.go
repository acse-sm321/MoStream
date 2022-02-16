package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// init a router and return
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}

// listen->RegisterHandlers->handlers
// Create thousands of goroutines on multiple cores

// handler->validate{requests,user}-> logical operations -> response. [API services layer]
// validation-> data structure and err handling
// session? what is session and how to use
