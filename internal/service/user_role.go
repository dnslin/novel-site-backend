package service

import (
    "context"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type UserRoleService interface {
	GetUserRole(ctx context.Context, id int64) (*model.UserRole, error)
}
func NewUserRoleService(
    service *Service,
    userRoleRepository repository.UserRoleRepository,
) UserRoleService {
	return &userRoleService{
		Service:        service,
		userRoleRepository: userRoleRepository,
	}
}

type userRoleService struct {
	*Service
	userRoleRepository repository.UserRoleRepository
}

func (s *userRoleService) GetUserRole(ctx context.Context, id int64) (*model.UserRole, error) {
	return s.userRoleRepository.GetUserRole(ctx, id)
}
