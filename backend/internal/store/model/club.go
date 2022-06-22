package model

func (Club) TableName() string {
	return "club"
}

type Club struct {
	Id          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	OwnerID     uint64 `gorm:"type:int" json:"owner_id"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:varchar(150)" json:"description"`
}
