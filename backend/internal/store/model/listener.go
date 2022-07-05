package model

func (Listener) TableName() string {
	return "listener"
}

type Listener struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID   uint64 `gorm:"type:datetime" json:"user_id"`
	ClubID   string `gorm:"type:int" json:"club_id"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
}
