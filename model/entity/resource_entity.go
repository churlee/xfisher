package entity

type ResourceEntity struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Type       string `json:"type" db:"type"`
	Content    string `json:"content" db:"content"`
	LinkUrl    string `json:"linkUrl" db:"link_url"`
	LinkPwd    string `json:"linkPwd" db:"link_pwd"`
	View       int    `json:"view" db:"view"`
	UserId     string `json:"userId" db:"user_id"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
