package dto

import "lilith/model/entity"

type UserDto struct {
	UserInfo entity.UserInfo `json:"userInfo"`
	Token    string          `json:"token"`
}

type AuthDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePwdDto struct {
	OPwd string `json:"oPwd" validate:"required"`
	NPwd string `json:"nPwd" validate:"required"`
	RPwd string `json:"rPwd" validate:"required,eqfield=NPwd"`
}

type UpdateUserDto struct {
	Nick     string `json:"nick" db:"nick" validate:"required"`
	Position string `json:"position" db:"position" validate:"required"`
}
