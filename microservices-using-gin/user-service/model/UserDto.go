package model

type UserDto struct {
	Id       int
	UserName string `json:"userName"`
	Password string `json:"password"`
}
