package model

type UserAddress struct {
	Id      int64  `gorm:"type:bigint;not null;primary_key;"`
	Name    string `gorm:"type:varchar(20);not null"`
	Phone   string `gorm:"type:varchar(20);not null"`
	Area    string `gorm:"type:varchar(50);not null"`
	Address string `gorm:"type:varchar(50);not null"`
	Company string `gorm:"type:varchar(50);not null"`
}

func (UserAddress) TableName() string {
	return "useraddress"
}
