package model

import (
	"time"

	"gorm.io/gorm"
)

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
}

type RatingType struct {
	Id          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	Description string
	Level       int `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type BookRating struct {
	Id           uint `gorm:"primarykey"`
	BookId       uint `gorm:"not null"`
	RatingTypeId uint `gorm:"not null"`
	Comment      string
	IP           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (b *Book) TableName() string {
	return "books"
}

func (rt *RatingType) TableName() string {
	return "rating_types"
}

func (br *BookRating) TableName() string {
	return "book_ratings"
}
