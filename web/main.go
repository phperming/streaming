package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/",HomeHandler)
	return router
}

func main() {
	r := RegisterHandler()
	fmt.Println("start")
	err := http.ListenAndServe(":9091",r)
	if err != nil {
		fmt.Println(err)
	}
}
