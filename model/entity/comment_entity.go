package entity

type CommentEntity struct {
	Id         string `json:"id" db:"id"`
	ResourceId string `json:"resourceId" db:"resource_id"`
	UserId     string `json:"userId" db:"user_id"`
	Content    string `json:"content" db:"content"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
