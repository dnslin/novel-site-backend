package model

type UserRole struct {
	Id     uint `gorm:"primarykey"`
	UserId uint `gorm:"not null;index" json:"user_id"` // 用户ID
	RoleId uint `gorm:"not null;index" json:"role_id"` // 角色ID
}

func (ur *UserRole) TableName() string {
	return "user_roles"
}
