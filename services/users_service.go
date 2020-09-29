package services

import (
	"github.com/nasruddin/golang/elastic/domain/users"
)

var (
	UserServiceInterface userServiceInterface = &userService{};
)

type userServiceInterface interface {
	Create(user users.User) (*users.User, error)
	Get(string) (*users.User, error)
}

type userService struct {}

func (s userService) Create(user users.User) (*users.User, error)  {
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s userService) Get(id string) (*users.User, error)  {
	user := users.User{ID: id}

	if err := user.Get(); err != nil {
		return nil, nil
	}

	return &user, nil
}