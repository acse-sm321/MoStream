package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type homePage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// implement session and cookie
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := homePage{
			Name: "Shuheng Mo",
		}

		t, e := template.ParseFiles("./templates/home.html")
		if e != nil {
			log.Printf("Parsing templates home.html error: %s", e)
			return
		}
		t.Execute(w, p) // add the variable to your templates
		return
	}

	// if user, direct to page
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

	// the authentication of session ID is implemented in the front-end
}

// auth: users and visitors
