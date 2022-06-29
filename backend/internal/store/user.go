package store

import (
	"backend/internal/store/model"
	"gorm.io/gorm"
)

type UserStore struct {
	connection *gorm.DB
}

func NewUserStore(connection *gorm.DB) *UserStore {
	return &UserStore{
		connection,
	}
}

func (u *UserStore) SaveUser(user model.User) *model.User {
	u.connection.Where(model.User{SpotifyID: user.SpotifyID}).FirstOrCreate(&user)
	return &user
}

func (u *UserStore) AddUserLogon(userLogon *model.UserLogon) *model.UserLogon {
	var user model.UserLogon
	u.connection.Where("user_id = ?", userLogon.UserID).Last(&user)
	if user.UserID != 0 {
		var user model.UserLogon
		u.connection.Where("user_id = ?", userLogon.UserID).Last(&user).Update("token", userLogon.Token)
	} else {
		u.connection.Save(userLogon)

	}
	return userLogon
}

func (u *UserStore) GetUserActiveUserLogon(userId uint64) *model.UserLogon {
	var user model.UserLogon
	u.connection.Where("user_id = ?", userId).Last(&user)
	return &user
}

func (u *UserStore) GetUserById(userId uint64) *model.User {
	var user model.User
	u.connection.First(&user, userId)
	return &user
}
