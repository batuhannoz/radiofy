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

func (c *ClubStore) ClubByOwnerID(ownerID uint64) *model.Club {
	var club model.Club
	c.connection.Where("owner_id = ?", ownerID).Last(&club)
	return &club
}

func (c *ClubStore) ClubListeners(ClubID int) *[]model.Listener {
	var listeners []model.Listener
	c.connection.Model(&listeners).Where("club_id = ?", ClubID).Where("is_active = ?", true).Scan(&listeners)
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

func (c *ClubStore) DeactivateListener(userID uint64) {
	c.connection.Model(&model.Listener{}).Where("user_id = ?", userID).UpdateColumn("is_active", false)
}
