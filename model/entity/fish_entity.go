package entity

type FishEntity struct {
	Id         string `db:"id" json:"id"`
	Title      string `db:"title" json:"title"`
	Url        string `db:"url" json:"url"`
	Order      int    `db:"order" json:"order"`
	Type       string `db:"type" json:"type"`
	View       string `db:"view" json:"view"`
	CreateTime int64  `db:"create_time" json:"createTime"`
}
