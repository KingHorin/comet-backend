package model

type Item struct {
	Id    int64  `gorm:"type:bigint;not null;primary_key;auto_increment"`
	Name  string `gorm:"type:varchar(50);not null;index"`
	Price int32  `gorm:"type:int;not null;"`
	Stock int32  `gorm:"type:int;not null"`
}

func (Item) TableName() string {
	return "items"
}
