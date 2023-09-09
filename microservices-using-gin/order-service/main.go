package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"order-service/dependencies"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//loadEnv()
	dependencies.InitializeDependencies()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/order")
	publicRoutes.POST("/save", dependencies.OrderControllerInstance.SaveOrder)

	//protectedRoutes := router.Group("/api")
	//protectedRoutes.Use(middleware.JWTAuthMiddleware())

	//protectedRoutes.GET("/user/:name", controller.GetUserDetailsByName)
	//protectedRoutes.GET("/user", controller.GetUserDetailsById)

	err := router.Run(":8001")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port 8001")
}
