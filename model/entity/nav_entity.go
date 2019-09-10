package entity

type NavEntity struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Type       string `json:"type" db:"type"`
	Desc       string `json:"desc" db:"desc"`
	Icon       string `json:"icon" db:"icon"`
	Like       string `json:"like" db:"like"`
	View       string `json:"view" db:"view"`
	Url        string `json:"url" db:"url"`
	Collection string `json:"collection" db:"collection"`
}
