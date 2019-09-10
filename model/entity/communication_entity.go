package entity

type CommunicationEntity struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
	Node       string `json:"node" db:"node"`
	UserId     string `json:"userId" db:"user_id"`
	View       int    `json:"view" db:"view"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
