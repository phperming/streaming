package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.GET("/video-delete-record/:vid-id",vidDelRecHandler)
	return r
}

func main() {
	r := RegisterHandlers()

	http.ListenAndServe("9001",r)
}
