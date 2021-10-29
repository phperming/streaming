package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)
var tempvid string
func clearTables()  {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add",testAddUser)
	t.Run("GetUser",testGetUser)
	t.Run("DeleteUser",testDeleteUser)
	t.Run("ReGetUser",testReGetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("Michle","123456")
	if err != nil {
		t.Errorf("Error of AddUser:%v",err)
	}
}

func testGetUser(t *testing.T)  {
	pwd,err := GetUserCredential("Michle")
	if err != nil || pwd != "123456" {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("Michle","123456")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v",err)
	}
}

func testReGetUser( t *testing.T) {
	pwd,err := GetUserCredential("Michle")
	if err != nil {
		t.Errorf("Error of RegetUser :%v",err)
	}
	if pwd != "" {
		t.Errorf("DeleteUser test Faild")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddVideo",testAddVideoInfo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DeleteVideo",testDeleteVideoInfo)
	t.Run("ReGetVideo",testReGetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi,err := AddNewVideo(1,"my-video")
	if err != nil {
		t.Errorf("Error of AddVideo:%v",err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T)  {
	_,err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v",err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v",err)
	}
}

func testReGetVideoInfo(t *testing.T)  {
	vi,err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of ReGetVideoInfo:%v",err)
	}

}

func TestCommentWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("AddUser",testAddUser)
	t.Run("AddComment",testAddComment)
	t.Run("ListComments",testListComments)

}

func testAddComment(t *testing.T)  {
	err := AddNewComment("123456",1,"is very good!")
	if err != nil {
		t.Errorf("Error of AddComment:%v",err)
	}
}

func testListComments(t *testing.T) {
	from := 1635230981
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res,err := ListComments("123456",from,to)
	if err != nil {
		t.Errorf("Error of ListComments:%v",err)
	}
	for i,ele := range res {
		fmt.Printf("comment:%d,%v",i,ele)
	}
}
