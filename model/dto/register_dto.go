package dto

type RegisterDto struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"rePassword" validate:"required,eqfield=Password"`
}
