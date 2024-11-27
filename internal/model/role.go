package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
}

func (m *Role) TableName() string {
    return "role"
}
