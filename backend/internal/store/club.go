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

func (c *ClubStore) Clubs() *model.Club {
	return nil
}

func (c *ClubStore) CreateClub(club *model.Club) *model.Club {
	c.connection.Save(&club)
	return club
}
