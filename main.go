package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thampaponn/learn-go/controller"
	"github.com/thampaponn/learn-go/database"
	"github.com/thampaponn/learn-go/service"
)

func main() {
	database.InitDB()
	server := gin.Default()

	userService := service.New()
	userController := controller.New(userService)

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(201, userController.Save(ctx))
	})

	server.Run(":8080")
}
