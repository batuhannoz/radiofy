package model

func (ClubSong) TableName() string {
	return "club_song"
}

type ClubSong struct {
	Id         int    `gorm:"primary_key:auto_increment" json:"id"`
	ClubID     int    `gorm:"type:int" json:"club_id"`
	AlbumID    string `gorm:"type:varchar(100)" json:"album_id"`
	Position   int    `gorm:"type:int" json:"position"`
	SongID     string `gorm:"varchar(100)" json:"song_id"`
	SongName   string `gorm:"varchar(100)" json:"song_name"`
	ArtistName string `gorm:"varchar(100)" json:"artist_name"`
	Image      string `gorm:"varchar(150)" json:"image"`
	ProgressMS int    `gorm:"type:int" json:"progress_ms"`
}
