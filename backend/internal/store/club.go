package store

import "gorm.io/gorm"

type ClubStore struct {
	connection *gorm.DB
}

func NewClubStore(connection *gorm.DB) *ClubStore {
	return &ClubStore{
		connection,
	}
}
