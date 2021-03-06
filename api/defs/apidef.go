package defs

type User struct {
	Id int `json:"id"`
	LoginName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id string `json:"id""`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string
	TTL int64
}


