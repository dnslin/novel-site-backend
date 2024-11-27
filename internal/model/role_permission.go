package model

type RolePermission struct {
	Id           uint `gorm:"primarykey"`
	RoleId       uint `gorm:"not null;index" json:"role_id"`       // 角色ID
	PermissionId uint `gorm:"not null;index" json:"permission_id"` // 权限ID
}

func (rp *RolePermission) TableName() string {
	return "role_permissions"
}
