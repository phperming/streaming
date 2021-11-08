package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}


func HomeHandler(w http.ResponseWriter,r *http.Request, ps httprouter.Params)  {
	cname,err := r.Cookie("username")
	sid,err2 := r.Cookie("session")
	if err != nil || err2 != nil {
		p := &HomePage{Name: "Michel"}
		t,e := template.ParseFiles("./templates/home.html")
		if e != nil {
			log.Printf("Parse file home.html failed: %s",err)
			return
		}
		t.Execute(w,p)
		return
	}

	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w,r,"/userhome",301)
		return
	}
}
