package model

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Id       int
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}
