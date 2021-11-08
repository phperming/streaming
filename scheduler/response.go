package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter,sc int,errMsg string)  {
	w.WriteHeader(sc)
	io.WriteString(w,errMsg)
}
