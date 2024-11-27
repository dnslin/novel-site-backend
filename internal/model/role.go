package model

import (
	"time"

	"gorm.io/gorm"
)

// Role 角色表
type Role struct {
	Id          uint         `gorm:"primarykey"`
	Name        string       `gorm:"size:50;not null;unique" json:"name"` // 角色名称
	Code        string       `gorm:"size:50;not null;unique" json:"code"` // 角色编码
	Description string       `gorm:"size:200" json:"description"`         // 角色描述
	Permissions []Permission `gorm:"many2many:role_permissions;"`         // 角色权限多对多关系
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (r *Role) TableName() string {
	return "roles"
}
