package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thampaponn/learn-go/entity"
	"github.com/thampaponn/learn-go/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return entity.Video{}
	}
	c.service.Save(video)
	return video
}
