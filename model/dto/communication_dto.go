package dto

type ListCommunicationDto struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	View       int    `json:"view" db:"view"`
	Nick       string `json:"nick" db:"nick"`
	Photo      string `json:"photo" db:"photo"`
	Position   string `json:"position" db:"position"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}

type CreateCommunicationDto struct {
	Title   string `json:"title" db:"title" validate:"required"`
	Node    string `json:"node" db:"node" validate:"required"`
	Content string `json:"content" db:"content" validate:"required"`
}

type CommunicationInfoDto struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Node       string `json:"node" db:"node"`
	Content    string `json:"content" db:"content"`
	View       int    `json:"view" db:"view"`
	Nick       string `json:"nick" db:"nick"`
	Photo      string `json:"photo" db:"photo"`
	Position   string `json:"position" db:"position"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
