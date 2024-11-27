package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string `gorm:"unique;not null"`
	Nickname  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Avatar    string `gorm:"comment:头像地址"`
	Intro     string `gorm:"comment:个人简介"`
	Email     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}
