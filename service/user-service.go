package service

import (
	"github.com/thampaponn/learn-go/database"
	"github.com/thampaponn/learn-go/entity"
)

type UserService interface {
	FindAll() []entity.User
	Save(entity.User) entity.User
}

type userService struct {
}

func New() UserService {
	return &userService{}
}

func (service *userService) FindAll() []entity.User {
	var users []entity.User
	database.DB.Find(&users)
	return users
}

func (service *userService) Save(user entity.User) entity.User {
	database.DB.Create(&user)
	return user
}
