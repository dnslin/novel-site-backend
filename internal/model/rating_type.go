package model

import (
	"time"

	"gorm.io/gorm"
)

// RatingType 评分类型实体
type RatingType struct {
	Id          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	Description string
	Level       int `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (rt *RatingType) TableName() string {
	return "rating_types"
}
