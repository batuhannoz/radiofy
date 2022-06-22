package model

func (ClubSong) TableName() string {
	return "club_song"
}

type ClubSong struct {
	Id     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ClubID uint64 `gorm:"type:int" json:"club_id"`
	SongID uint64 `gorm:"type:varchar(50)" json:"song_id"`
}
