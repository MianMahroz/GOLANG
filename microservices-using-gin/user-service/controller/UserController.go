package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/model"
	"user-service/service"
)

func Register(context *gin.Context) {
	var input model.UserDto

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserServiceImpl{UserDto: input}
	savedUser, err := userService.RegisterUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"msg": savedUser})
}

func Login(context *gin.Context) {
	var input model.UserDto

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserServiceImpl{UserDto: input}
	savedUser, err := userService.LoginUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"msg": savedUser})
}

func GetUserDetailsByName(context *gin.Context) {

	var userName = context.Param("name")
	if len(userName) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Param 'name' not found"})
		return
	}

	dto := model.UserDto{UserName: userName}
	user := service.UserServiceImpl{UserDto: dto}
	userDto, err := user.GetUserByName()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": userDto})
}
