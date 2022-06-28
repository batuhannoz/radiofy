package store

import (
	"backend/internal/store/model"
	"gorm.io/gorm"
)

type ClubStore struct {
	connection *gorm.DB
}

func NewClubStore(connection *gorm.DB) *ClubStore {
	return &ClubStore{
		connection,
	}
}

func (club *ClubStore) Clubs() *model.Club {
	return nil
}
