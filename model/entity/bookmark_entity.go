package entity

type BookmarkEntity struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Url        string `json:"url" db:"url"`
	UserId     string `json:"userId" db:"user_id"`
	Share      int    `json:"share" db:"share"`
	Tags       string `json:"tags" db:"tags"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}

type BookmarkTagEntity struct {
	Id         string `json:"id" db:"id"`
	BookmarkId string `json:"bookmarkId" db:"bookmark_id"`
	Tag        string `json:"tag" db:"tag"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
