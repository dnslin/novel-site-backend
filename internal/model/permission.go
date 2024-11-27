package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
}

func (m *Permission) TableName() string {
    return "permission"
}
