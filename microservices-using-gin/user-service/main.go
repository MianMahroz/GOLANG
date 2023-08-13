package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"user-service/controller"
	"user-service/database"
	"user-service/middleware"
	"user-service/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	err := database.Database.AutoMigrate(&model.UserEntity{})
	if err != nil {
		return
	}

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.GET("/user/:name", controller.GetUserDetailsByName)
	protectedRoutes.GET("/user", controller.GetUserDetailsById)

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port 8000")
}
