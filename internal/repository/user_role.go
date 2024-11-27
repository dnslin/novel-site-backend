package repository

import (
    "context"
	"novel-site-backend/internal/model"
)

type UserRoleRepository interface {
	GetUserRole(ctx context.Context, id int64) (*model.UserRole, error)
}

func NewUserRoleRepository(
	repository *Repository,
) UserRoleRepository {
	return &userRoleRepository{
		Repository: repository,
	}
}

type userRoleRepository struct {
	*Repository
}

func (r *userRoleRepository) GetUserRole(ctx context.Context, id int64) (*model.UserRole, error) {
	var userRole model.UserRole

	return &userRole, nil
}
