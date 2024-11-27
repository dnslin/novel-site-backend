package model

import (
	"time"

	"gorm.io/gorm"
)

// BookRating 书籍评分实体
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

// RatingTypeCount 评分类型统计
type RatingTypeCount struct {
	RatingTypeID uint  `json:"rating_type_id"`
	Count        int64 `json:"count"`
}

func (br *BookRating) TableName() string {
	return "book_ratings"
}
