package service

import (
    "context"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type RolePermissionService interface {
	GetRolePermission(ctx context.Context, id int64) (*model.RolePermission, error)
}
func NewRolePermissionService(
    service *Service,
    rolePermissionRepository repository.RolePermissionRepository,
) RolePermissionService {
	return &rolePermissionService{
		Service:        service,
		rolePermissionRepository: rolePermissionRepository,
	}
}

type rolePermissionService struct {
	*Service
	rolePermissionRepository repository.RolePermissionRepository
}

func (s *rolePermissionService) GetRolePermission(ctx context.Context, id int64) (*model.RolePermission, error) {
	return s.rolePermissionRepository.GetRolePermission(ctx, id)
}
