package model

import "time"

func (User) TableName() string {
	return "user"
}

type User struct {
	Id          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	DisplayName string    `gorm:"type:varchar(100)" json:"display_name"`
	Image       string    `gorm:"type:varchar(150)" json:"image"`
	CreateDate  time.Time `gorm:"type:datetime" json:"create_date"`
	SpotifyID   string    `gorm:"type:varchar(50)" json:"spotify_id"`
	Mail        string    `gorm:"type:varchar(100)" json:"mail"`
	Country     string    `gorm:"type:varchar(4)" json:"country"`
	Product     string    `gorm:"type:varchar(10)" json:"product"`
}
