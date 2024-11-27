package model

import (
	"time"

	"gorm.io/gorm"
)

// Permission 权限表
type Permission struct {
	Id          uint   `gorm:"primarykey"`
	Name        string `gorm:"size:50;not null" json:"name"`        // 权限名称
	Code        string `gorm:"size:50;not null;unique" json:"code"` // 权限编码
	Type        string `gorm:"size:20;not null" json:"type"`        // 权限类型(menu,button,api)
	ParentId    uint   `gorm:"default:0" json:"parent_id"`          // 父级ID
	Path        string `gorm:"size:200" json:"path"`                // 路径
	Description string `gorm:"size:200" json:"description"`         // 描述
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (p *Permission) TableName() string {
	return "permissions"
}
