package repo

import (
	"user-service/database"
	"user-service/model"
)

type UserRepo struct {
	Entity model.UserEntity
}

func (repo UserRepo) CreateUser() (string, error) {
	err := database.Database.Create(&repo.Entity).Error
	if err != nil {
		return "SOMETHING WENT WRONG!", err
	}

	return "CREATED", nil
}

func (repo UserRepo) FindUserByName() (model.UserEntity, error) {

	database.Database.Where("userName=?", repo.Entity.UserName).Find(&repo.Entity)

	return repo.Entity, nil
}

func (repo UserRepo) FindUserById() (model.UserEntity, error) {

	database.Database.First(&repo.Entity, repo.Entity.Id)
	return repo.Entity, nil
}
