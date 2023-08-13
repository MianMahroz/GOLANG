package model

type UserDto struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}
