package main

import (
	"MoStream/api/session"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
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

	// login methods
	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)
	router.GET("/user/:username", GetUserInfo)

	// videos and comments
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

// load session data from database
func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh) // apiserver runs on :8000
}

// listen->RegisterHandlers->handlers
// Create thousands of goroutines on multiple cores

// handler->validate{requests,user}-> logical operations -> response. [API services layer]
// validation-> data structure and err handling
// session? what is session and how to use
// middleware defs(message,err), handlers dbops and response
