package session

import (
	"sync"
	"time"
	"vedio_server/api/dbops"
	"vedio_server/api/defs"
	"vedio_server/api/utils"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNewSessionId(un string) string {
	id,_ := utils.NewUUID()
	ct := nowInMilli()

	ttl := ct + 30 * 60 * 1000
	ss := &defs.SimpleSession{Username: un,TTL: ttl}
	//保存到sessionMap中
	sessionMap.Store(id,ss)
	//写入数据库
	err := dbops.InsertSession(id,ttl,un)

	if err != nil {
		return ""
	}
	return id
}

func LoadSessionFromDB()  {
	r,err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

func IsSessionExpired(sid string) (string ,bool)  {
	ss,ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "",true
		}

		return ss.(*defs.SimpleSession).Username,false
	} else {
		ss,err := dbops.RetrieveSessions(sid)
		if err != nil || ss == nil {
			return "",true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "",true
		}

		sessionMap.Store(sid,ss)

		return ss.Username,false
	}

	return "" , true
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}


