package repository

import (
    "context"
	"novel-site-backend/internal/model"
)

type RolePermissionRepository interface {
	GetRolePermission(ctx context.Context, id int64) (*model.RolePermission, error)
}

func NewRolePermissionRepository(
	repository *Repository,
) RolePermissionRepository {
	return &rolePermissionRepository{
		Repository: repository,
	}
}

type rolePermissionRepository struct {
	*Repository
}

func (r *rolePermissionRepository) GetRolePermission(ctx context.Context, id int64) (*model.RolePermission, error) {
	var rolePermission model.RolePermission

	return &rolePermission, nil
}
