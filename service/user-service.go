package service

import (
	"errors"

	"github.com/thampaponn/learn-go/database"
	"github.com/thampaponn/learn-go/entity"
)

type UserService interface {
	FindAll() []entity.User
	Save(entity.User) (entity.User, error)
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

func (service *userService) Save(user entity.User) (entity.User, error) {
	// Check required fields
	if user.FirstName == "" || user.LastName == "" || user.Username == "" || user.Password == "" {
		return entity.User{}, errors.New("missing required field(s)")
	}

	// Check for duplicate username
	var existing entity.User
	if err := database.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		return entity.User{}, errors.New("username already exists")
	}

	// Save user
	if err := database.DB.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
