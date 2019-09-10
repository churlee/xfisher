package entity

type FeedbackEntity struct {
	Id         string `json:"id" db:"id"`
	UserId     string `json:"userId" db:"user_id"`
	Content    string `json:"content" db:"content"`
	CreateTime int64  `json:"createTime" db:"create_time"`
}
