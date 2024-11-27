package repository

import (
    "context"
	"novel-site-backend/internal/model"
)

type PermissionRepository interface {
	GetPermission(ctx context.Context, id int64) (*model.Permission, error)
}

func NewPermissionRepository(
	repository *Repository,
) PermissionRepository {
	return &permissionRepository{
		Repository: repository,
	}
}

type permissionRepository struct {
	*Repository
}

func (r *permissionRepository) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	var permission model.Permission

	return &permission, nil
}
