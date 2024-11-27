package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string `gorm:"unique;not null"`
	Nickname  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Avatar    string `gorm:"comment:头像地址"`
	Intro     string `gorm:"comment:个人简介"`
	Email     string `gorm:"not null"`
	Status    int    `gorm:"default:1" json:"status"` // 状态 1:正常 2:禁用
	Roles     []Role `gorm:"many2many:user_roles"`    // 用户角色多对多关系
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

// HasPermission 检查用户是否有某个权限
func (u *User) HasPermission(permissionCode string) bool {
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if perm.Code == permissionCode {
				return true
			}
		}
	}
	return false
}
