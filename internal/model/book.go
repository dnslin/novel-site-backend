package model

import (
	"time"

	"gorm.io/gorm"
)

// Book 书籍实体
type Book struct {
	Id          uint   `gorm:"primarykey"`
	FileName    string `gorm:"column:file_name;not null"`
	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	FileSize    int64  `gorm:"column:file_size;not null"`
	MD5         string `gorm:"column:md5;unique;not null"`
	NewFileName string `gorm:"column:new_file_name;not null"`
	Cover       string
	Intro       string
	Parts       string
	FileURL     string `gorm:"column:file_url"`
	Sort        string
	Type        string
	Tag         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	HotValue    int64          `gorm:"column:hot_value;default:0"`
	Downloads   int64          `gorm:"column:downloads;default:0"`
}

func (b *Book) TableName() string {
	return "books"
}
