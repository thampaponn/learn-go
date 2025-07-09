package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thampaponn/learn-go/controller"
	"github.com/thampaponn/learn-go/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(":8080")
}
