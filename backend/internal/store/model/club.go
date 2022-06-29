package model

func (Club) TableName() string {
	return "club"
}

type Club struct {
	Id          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	OwnerID     uint64 `gorm:"type:int" json:"owner_id"`
	ClubCode    string `gorm:"type:varchar(6)" json:"club_code"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:varchar(150)" json:"description"`
	Image       string `gorm:"type:varchar(100)" json:"image"`
}
