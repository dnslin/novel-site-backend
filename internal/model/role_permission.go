package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
}

func (m *RolePermission) TableName() string {
    return "role_permission"
}
