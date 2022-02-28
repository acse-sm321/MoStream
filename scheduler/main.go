package main

import (
	"MoStream/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	//c := make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	//<- c
	http.ListenAndServe(":9001", r) // this is blocking
}
