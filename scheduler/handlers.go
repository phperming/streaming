package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"vedio_server/scheduler/dbops"
)

func vidDelRecHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
	vid :=  ps.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w,400,"video id should not be empty")
		return
	}
	
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w,500,"Internal Error")
		return
	}

	sendResponse(w,200, "")
	return

}