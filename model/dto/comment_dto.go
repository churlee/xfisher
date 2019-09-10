package dto

type CreateCommentDto struct {
	ResourceId string `json:"resourceId" validate:"required"`
	Content    string `json:"content" validate:"required"`
}

type CommentListDto struct {
	Id         string `json:"id" db:"id"`
	Content    string `json:"content" db:"content"`
	CreateTime int64  `json:"createTime" db:"create_time"`
	Nick       string `json:"nick" db:"nick"`
	Photo      string `json:"photo" db:"photo"`
	Position   string `json:"position" db:"position"`
}
