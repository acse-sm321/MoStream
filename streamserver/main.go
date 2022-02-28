package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter // flow control
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// get the video id
	router.GET("/videos/:vid-id", streamHandler)

	// upload with new id in handler
	router.POST("/upload/:vid-id", uploadHandler)

	return router
}

// handlers

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2) // ... connections in total
	http.ListenAndServe(":9000", mh)
}
