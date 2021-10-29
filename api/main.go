package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterRouter()  *httprouter.Router {
	r := httprouter.New()

	r.POST("/user",CreateUser)
	r.POST("/user/:username",Login)
	return r

}

func main() {
	r := RegisterRouter()

	http.ListenAndServe(":8000",r)
}
