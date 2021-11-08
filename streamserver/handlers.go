package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func testHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	t,_ := template.ParseFiles("./videos/upload.html")
	t.Execute(w,nil)
}

func streamHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	log.Println("Entered the StreamHandler")
	targetUrl := "http://127.0.0.1:9090/videos/" + ps.ByName("vid-id")

	http.Redirect(w,r,targetUrl,301)
}

func uploadHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);err != nil {
		sendErrorResponse(w,http.StatusBadRequest,"File is too big")
		return
	}

	file,ext,err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v ",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}
	fileSuffix := path.Ext(ext.Filename)
	data ,err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}
	fn := ps.ByName("vid-id") + fileSuffix
	err = ioutil.WriteFile(VIDEO_DIR + fn,data, 666)

	if err != nil {
		log.Printf("Write File error: %v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"Uploaded Successfully")

}
