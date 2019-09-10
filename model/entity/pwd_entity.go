package entity

type UserPwd struct {
	Id         string `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	Salt       string `json:"salt" db:"salt"`
	IsOK       string `json:"isOk" db:"is_ok"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
