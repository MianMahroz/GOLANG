package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"order-service/dependencies"
)

func main() {
	dependencies.InitializeDependencies()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/saveOrder", dependencies.OrderControllerInstance.SaveOrder)

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
