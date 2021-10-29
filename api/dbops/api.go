package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"vedio_server/api/defs"
	"vedio_server/api/utils"
)

func AddUserCredential(loginName string,pwd string) error {
	stmtIns , err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?, ?)")

	if err != nil {
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error)  {
	stmtOut ,err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		return "",err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error {
	stmtDelete, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_,err = stmtDelete.Exec(loginName,pwd)

	if err != nil {
		return err
	}
	defer stmtDelete.Close()
	return nil
}

func GetUser(loginName string) (*defs.User,error) {
	stmtOut,err := dbConn.Prepare("SELECT id,pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s",err)
		return nil,err
	}

	var id int
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&id,&pwd)
	if err != nil {
		return nil,err
	}
	if err == sql.ErrNoRows {
		return nil,nil
	}
	res := &defs.User{Id:id,LoginName: loginName,Pwd: pwd}
	defer stmtOut.Close()
	return res, nil
}

func AddNewVideo(aid int,name string) (*defs.VideoInfo,error)  {
	vid,err := utils.NewUUID()
	if err != nil {
		return nil,err
	}
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04:05")

	stmtIns,err := dbConn.Prepare("INSERT  INTO video_info (id,author_id,name,display_ctime) VALUES (?,?,?,?)")
	if err != nil {
		return nil,err
	}

	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid,AuthorId: aid,Name: name,DisplayCtime: ctime}
	stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo,error) {
	stmtOut,err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id=?")
	if err != nil {
		return nil, err
	}

	var author_id int
	var name,display_ctime string

	err = stmtOut.QueryRow(vid).Scan(&author_id,&name,&display_ctime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("getVedio Error:%v",err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{
		Id: vid,
		AuthorId: author_id,
		Name: name,
		DisplayCtime: display_ctime,
	}

	return res, nil
}

func ListVideoInfo(uname string,from,to int) ([]*defs.VideoInfo,error) {
	stmtOut,err := dbConn.Prepare(`SELECT video_info.id,video_info.author_id,video_info.name,video_info.display_ctime FROM video_info 
		INNER JOIN users ON users.id=video_info.author_id WHERE users.login_name=? AND video_info.create_time > FROM_UNIXTIME(?)
		AND video_info.create_time <= FROM_UNIXTIME(?) ORDER BY video_info.create_time DESC`)
	if err !=nil {
		return nil, err
	}
	var res  []*defs.VideoInfo

	rows,err := stmtOut.Query(uname,from,to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var vid,name,display_ctime string
		var author_id int
		if err := rows.Scan(&vid,&author_id,&name,&display_ctime);err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{
			Id: vid,
			AuthorId: author_id,
			Name: name,
			DisplayCtime: display_ctime,
		}

		res = append(res, vi)
	}
	defer stmtOut.Close()
	return res,nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel ,err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_,err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func AddNewComment(vid string,aid int,content string) error {
	id,err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns ,err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}
	_,err =stmtIns.Exec(id,vid,aid,content)

	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string,from,to int) ([]*defs.Comment,error) {
	stmtOut,err := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content FROM comments INNER JOIN users ON users.id=comments.author_id
					WHERE comments.video_id=? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?) ORDER BY comments.time DESC`)
	if err != nil {
		return nil, err
	}

	rows,err := stmtOut.Query(vid,from,to)

	if err != nil {
		return nil, err
	}
	var res []*defs.Comment
	for rows.Next() {
		var cid,login_name ,content string
		if err := rows.Scan(&cid,&login_name,&content);err != nil {
			return res, err
		}
		cm := &defs.Comment{
			Id: cid,
			VideoId: vid,
			Author: login_name,
			Content: content,
		}
		res = append(res, cm)
	}
	defer stmtOut.Close()
	return res,nil
}