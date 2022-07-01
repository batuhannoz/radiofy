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

func (c *ClubStore) CreateClub(club *model.Club) *model.Club {
	c.connection.Save(&club)
	return club
}

func (c *ClubStore) Clubs() *[]model.Club {
	var clubs []model.Club
	c.connection.Raw("SELECT * FROM club WHERE is_active=true").Scan(&clubs)
	return &clubs
}
