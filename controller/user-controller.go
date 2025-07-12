package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thampaponn/learn-go/entity"
	"github.com/thampaponn/learn-go/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.User {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return entity.User{}
	}
	user, err := c.service.Save(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return entity.User{}
	}
	return user
}
