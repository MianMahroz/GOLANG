package service

import (
	"user-service/helper"
	"user-service/model"
	"user-service/repo"
)

type UserServiceImpl struct {
	UserDto model.UserDto
}

func (userService UserServiceImpl) RegisterUser() (string, error) {

	hashPass, _ := helper.Encode(userService.UserDto.Password) // encoding password
	userService.UserDto.UserName = helper.TrimString(userService.UserDto.UserName)

	var req = model.UserEntity{UserName: userService.UserDto.UserName, Password: hashPass}
	var userRepo = repo.UserRepo{Entity: req} // linking user with repo
	msg, err := userRepo.CreateUser()

	if err != nil {
		return msg, err
	}

	return msg, nil
}

func (userService UserServiceImpl) LoginUser() (string, error) {

	var userRepo = repo.UserRepo{Entity: model.UserEntity{UserName: userService.UserDto.UserName}} // linking user with repo

	user, err := userRepo.FindUserByName()

	if err != nil {
		return "NO USER FOUND!", err
	}

	msg, err := helper.VerifyPassword(user.Password, userService.UserDto.Password)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

func (userService UserServiceImpl) GetUserByName() (model.UserDto, error) {
	var userRepo = repo.UserRepo{Entity: model.UserEntity{UserName: userService.UserDto.UserName}}

	user, err := userRepo.FindUserByName()

	if err != nil {
		return model.UserDto{}, err
	}
	return model.UserDto{Id: user.Id, UserName: user.UserName}, nil
}
func (userService UserServiceImpl) GetUserById() (model.UserDto, error) {
	var userRepo = repo.UserRepo{Entity: model.UserEntity{Id: userService.UserDto.Id}}

	user, err := userRepo.FindUserById()

	if err != nil {
		return model.UserDto{}, err
	}
	return model.UserDto{Id: user.Id, UserName: user.UserName}, nil
}
