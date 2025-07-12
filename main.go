package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thampaponn/learn-go/controller"
	"github.com/thampaponn/learn-go/initializers"
	"github.com/thampaponn/learn-go/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.SyncDB()
}

func main() {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "It's running")
	})

	//User
	app.POST("/signup", controller.SignUp)
	app.POST("/login", controller.Login)
	app.GET("/validate", middleware.RequireAuth, controller.Validate)
	app.DELETE("/users/:id", controller.DeleteUser)

	app.Run()
}
