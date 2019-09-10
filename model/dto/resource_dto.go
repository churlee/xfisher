package dto

type CreateResourceDto struct {
	Title   string `json:"title" db:"title" validate:"required"`
	Type    string `json:"type" db:"type" validate:"required"`
	Content string `json:"content" db:"content" validate:"required"`
	LinkUrl string `json:"linkUrl" db:"link_url" validate:"required"`
	LinkPwd string `json:"linkPwd" db:"link_pwd"`
}

type ResourceInfoDto struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Type       string `json:"type" db:"type"`
	Content    string `json:"content" db:"content"`
	View       int    `json:"view" db:"view"`
	LinkUrl    string `json:"linkUrl" db:"link_url"`
	LinkPwd    string `json:"linkPwd" db:"link_pwd"`
	Nick       string `json:"nick" db:"nick"`
	Photo      string `json:"photo" db:"photo"`
	Position   string `json:"position" db:"position"`
	IsComment  bool   `json:"isComment"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}

type ResourceListDto struct {
	Id         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Type       string `json:"type" db:"type"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	Nick       string `json:"nick" db:"nick"`
}

type ResourceLinkAndPwdDto struct {
	LinkUrl string `json:"linkUrl" db:"link_url"`
	LinkPwd string `json:"linkPwd" db:"link_pwd"`
}
