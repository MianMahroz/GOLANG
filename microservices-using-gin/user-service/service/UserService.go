package service

type UserService interface {
	RegisterUser() string
	LoginUser() string
	GetUserByName() string
	GetUserBYId() string
}
