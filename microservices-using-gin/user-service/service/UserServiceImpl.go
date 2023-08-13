package service

import (
	"errors"
	"user-service/helper"
	"user-service/model"
	"user-service/repo"
)

type UserServiceImpl struct {
	UserDto model.UserDto
}

func (userService UserServiceImpl) RegisterUser() (string, error) {

	hashPass, _ := helper.Encode(userService.UserDto.Password) // encoding password
	userService.UserDto.Username = helper.TrimString(userService.UserDto.Username)

	var req = model.UserEntity{Username: userService.UserDto.Username, Password: hashPass}
	var userRepo = repo.UserRepo{Entity: req} // linking user with repo
	msg, err := userRepo.CreateUser()

	if err != nil {
		return msg, err
	}

	return msg, nil
}

func (userService UserServiceImpl) LoginUser() (model.UserDto, error) {

	var userRepo = repo.UserRepo{Entity: model.UserEntity{Username: userService.UserDto.Username}} // linking user with repo

	user, err := userRepo.FindUserByName()
	if err != nil || user.Id <= 0 {
		return model.UserDto{}, errors.New("no USER FOUND")
	}

	_, err = helper.VerifyPassword(user.Password, userService.UserDto.Password)
	if err != nil {
		return model.UserDto{}, errors.New("invalid password")
	}

	jwt, err := helper.GenerateJWT(user.Id)
	if err != nil {
		return model.UserDto{}, err
	}

	return model.UserDto{Id: user.Id, Username: user.Username, Password: user.Password, Token: jwt}, nil
}

func (userService UserServiceImpl) GetUserByName() (model.UserDto, error) {
	var userRepo = repo.UserRepo{Entity: model.UserEntity{Username: userService.UserDto.Username}}

	user, err := userRepo.FindUserByName()

	if err != nil || user.Id <= 0 {
		return model.UserDto{}, errors.New("no user found")
	}
	return model.UserDto{Id: user.Id, Username: user.Username, Password: user.Password}, nil
}

func (userService UserServiceImpl) GetUserById() (model.UserDto, error) {
	var userRepo = repo.UserRepo{Entity: model.UserEntity{Id: userService.UserDto.Id}}

	user, err := userRepo.FindUserById()

	if err != nil {
		return model.UserDto{}, err
	}
	return model.UserDto{Id: user.Id, Username: user.Username, Password: user.Password}, nil
}
