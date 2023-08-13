package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	dto := model.UserDto{Username: userName}
	user := service.UserServiceImpl{UserDto: dto}
	userDto, err := user.GetUserByName()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": userDto})
}

func GetUserDetailsById(context *gin.Context) {

	id := context.Query("id")
	if len(id) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Param 'id' not found"})
		return
	}

	val, err := strconv.Atoi(id)
	dto := model.UserDto{Id: val}
	user := service.UserServiceImpl{UserDto: dto}
	userDto, err := user.GetUserById()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": userDto})
}
