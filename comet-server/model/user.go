package model

type User struct {
	Id       int64  `gorm:"type:bigint;not null;primary_key;auto_increment"`
	Username string `gorm:"type:varchar(20);not null;index"`
	Password string `gorm:"type:varchar(20);not null;"`
	Phone    string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}
