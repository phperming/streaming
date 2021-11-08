package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"vedio_server/api/defs"
)

func InsertSession(id string,ttl int64,un string) error {
	ttlstr := strconv.FormatInt(ttl,10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_,err = stmtIns.Exec(id,ttlstr,un)
	if err != nil {
		return err
	}
	stmtIns.Close()

	return nil
}

func RetrieveSessions(id string) (*defs.SimpleSession,error)  {
	ss := &defs.SimpleSession{}
	stmtOut,err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")

	if err != nil {
		return nil, err
	}	
	var ttlstr string
	var login_name string
	err = stmtOut.QueryRow(id).Scan(&ttlstr,&login_name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if ttl, err := strconv.ParseInt(ttlstr,10,64);err == nil {
		ss.TTL = ttl
		ss.Username = login_name
	} else {
		return nil, err
	}
	defer stmtOut.Close()

	return ss, nil
}

func RetrieveAllSessions() (*sync.Map ,error ) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
		}
	}
	defer stmtOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel,err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		log.Printf("%s",err)
	}

	if _,err = stmtDel.Exec(sid);err != nil {
		return err
	}
	return nil
}