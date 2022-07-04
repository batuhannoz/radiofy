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

func (c *ClubStore) ClubByOwnerID(ownerID int) *model.Club {
	var club model.Club
	c.connection.Where("owner_id = ?", ownerID).Last(&club)
	return &club
}

func (c *ClubStore) ClubListeners(ClubID int) *[]model.Listener {
	var listeners []model.Listener
	c.connection.Raw("SELECT * FROM listener WHERE club_id = ?", ClubID).Scan(&listeners)
	return &listeners
}

func (c *ClubStore) AddListener(listener *model.Listener) *model.Listener {
	c.connection.Save(&listener)
	return listener
}

func (c *ClubStore) CurrentSongByClubID(clubID int) *model.ClubSong {
	var clubSong model.ClubSong
	c.connection.Where("club_id = ?", clubID).Last(&clubSong)
	return &clubSong
}

func (c *ClubStore) ChangeClubSong(clubSong *model.ClubSong) *model.ClubSong {
	c.connection.Save(&clubSong)
	return clubSong
}
