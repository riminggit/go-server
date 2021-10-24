package userModel

import (
	"time"
)

type User struct {
	ID        int    `gorm:"autoIncrement;primaryKey;type:int(11);not null"` // 自增
	Account   string `gorm:"type:varchar(50);uniqueIndex;not null"`          // 唯一索引
	NickName  string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(12)"`
	AvatarUrl string `gorm:"type:varchar(255)"`
	ComeFrom  string `gorm:"type:varchar(20)"`
	Openid    string `gorm:"type:varchar(50)"`
	Gender    int    `gorm:"type:int(3)"`
	City      string `gorm:"type:varchar(50)"`
	Province  string `gorm:"type:varchar(50)"`
	Country   string `gorm:"type:varchar(50)"`
	Language  string `gorm:"type:varchar(30)"`
	Rawdata   string `gorm:"type:varchar(255)"`
	Signature string `gorm:"type:varchar(255)"`
	IsAdmin   int    `gorm:"type:int(3)"`
	Birthday  time.Time
	CreateAt  time.Time
	DeleteAt  time.Time
	UpdateAt  time.Time
	IsUse     int `gorm:"type:int(3);default:1"`
}
