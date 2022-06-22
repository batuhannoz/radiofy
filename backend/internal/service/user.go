package service

import (
	"backend/internal/store/model"
)

type UserStore interface {
	Login(request *model.User) *model.User
}

type UserService struct {
	UserStore
}

func NewUserService(userStore UserStore) *UserService {
	return &UserService{
		userStore,
	}
}

func (u *UserService) Login(request model.User) *model.User {
	return u.UserStore.Login(&request)
}
