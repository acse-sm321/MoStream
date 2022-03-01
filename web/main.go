package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)

	// Proxy and transfer request
	router.POST("/api", apiHandler)

	router.POST("/upload/:vid-id", proxyHandler)

	// file server
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

// cross-origin resource sharing
// localhost:8080/upload/:vid-id
// localhost:9000/upload/:vid-id
