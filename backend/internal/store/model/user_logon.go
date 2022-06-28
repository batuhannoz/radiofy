package model

import "time"

func (UserLogon) TableName() string {
	return "user_logon"
}

type UserLogon struct {
	Id         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID     uint64    `gorm:"type:int" json:"user_id"`
	CreateDate time.Time `gorm:"type:datetime" json:"create_date"`
	Token      string    `gorm:"type:varchar(500)" json:"token"`
	IsLogout   bool      `gorm:"type:bool" json:"is_logout"`
}
