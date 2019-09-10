package dto

type PageDto struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
	Pages int         `json:"pages"`
}
