package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
}

func (m *UserRole) TableName() string {
    return "user_role"
}
