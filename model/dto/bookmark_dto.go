package dto

type CreateBookmarkDto struct {
	Title string   `json:"title" validate:"required"`
	Url   string   `json:"url" validate:"required"`
	Share string   `json:"share" validate:"required"`
	Tags  []string `json:"tags"`
}

type BookmarkListDto struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Url        string `json:"url" db:"url"`
	Nick       string `json:"nick" db:"nick"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	Tags       string `json:"tags" db:"tags"`
}
