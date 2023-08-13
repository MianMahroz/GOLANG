package repo

import (
	"log"
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

	err := database.Database.Where("Username=?", repo.Entity.Username).Find(&repo.Entity).Error
	if err != nil {
		return model.UserEntity{}, err
	}
	return repo.Entity, nil
}

func (repo UserRepo) FindUserById() (model.UserEntity, error) {

	log.Println("Id", repo.Entity.Id)

	err := database.Database.Where("id=?", repo.Entity.Id).Find(&repo.Entity).Error
	if err != nil {
		return model.UserEntity{}, err
	}

	return repo.Entity, nil
}
