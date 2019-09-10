package entity

type UserInfo struct {
	Id         string `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Nick       string `json:"nick" db:"nick"`
	Email      string `json:"email" db:"email"`
	Photo      string `json:"photo" db:"photo"`
	Desc       string `json:"desc" db:"desc"`
	Position   string `json:"position" db:"position"`
	HomePage   string `json:"homePage" db:"home_page"`
	Tags       string `json:"tags" db:"tags"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
