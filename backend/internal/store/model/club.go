package model

func (Club) TableName() string {
	return "club"
}

type Club struct {
	Id          int    `gorm:"primary_key:auto_increment" json:"id"`
	OwnerID     uint64 `gorm:"type:int" json:"owner_id"`
	IsActive    bool   `gorm:"type:bool" json:"is_active"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:varchar(150)" json:"description"`
	Image       string `gorm:"type:varchar(400)" json:"image"`
}
