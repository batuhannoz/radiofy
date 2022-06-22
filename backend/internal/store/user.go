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

func (db *UserStore) Login(request *model.User) *model.User {
	db.connection.Save(request)
	return request
}
