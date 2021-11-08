package taskrunner

import (
	"errors"
	"log"
	"sync"
	"vedio_server/scheduler/dbops"
)

func deleteVideo(vid string) error  {
	err := dbops.DelVideoDeletionRecord(vid)

	if err != nil {
		return err
	}

	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("video clear dispatcher error : %v",err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All task finished")
	}

	for _,id := range res {
		dc <- id
	}
	
	return nil
}

func VideoClearExecutor(dc dataChan) error  {
	errMap := &sync.Map{}
	var err error
	forloop:
		for {
			select {
			case vid := <- dc :
				go func(id interface{}) {
					if err := deleteVideo(id.(string));err != nil {
						errMap.Store(id,err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string));err != nil {
						errMap.Store(id,err)
					}
				}(vid)
			default:
				break forloop
			}
		}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}