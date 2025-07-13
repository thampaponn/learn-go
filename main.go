package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thampaponn/learn-go/controller"
	_ "github.com/thampaponn/learn-go/docs"
	"github.com/thampaponn/learn-go/initializers"
	"github.com/thampaponn/learn-go/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.SyncDB()
}

// @title Golang API
// @version 1.0
// @description This is a sample API for user management.

// @host localhost:3000
// @BasePath /api/v1
func main() {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "It's running")
	})

	// Swagger documentation

	v1 := app.Group("/api/v1")
	user := v1.Group("/users")

	{
		user.GET("/validate", middleware.RequireAuth, controller.Validate) // This will validate the user session
		user.POST("/signup", controller.SignUp)                            // This will create a new user
		user.POST("/login", controller.Login)                              // This will log in the user
		user.DELETE("/:id", middleware.RequireAuth, controller.DeleteUser) // This will delete the user with the given ID
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run()
}
