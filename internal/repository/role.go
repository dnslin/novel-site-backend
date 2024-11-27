package repository

import (
    "context"
	"novel-site-backend/internal/model"
)

type RoleRepository interface {
	GetRole(ctx context.Context, id int64) (*model.Role, error)
}

func NewRoleRepository(
	repository *Repository,
) RoleRepository {
	return &roleRepository{
		Repository: repository,
	}
}

type roleRepository struct {
	*Repository
}

func (r *roleRepository) GetRole(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role

	return &role, nil
}
