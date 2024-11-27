package service

import (
    "context"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type PermissionService interface {
	GetPermission(ctx context.Context, id int64) (*model.Permission, error)
}
func NewPermissionService(
    service *Service,
    permissionRepository repository.PermissionRepository,
) PermissionService {
	return &permissionService{
		Service:        service,
		permissionRepository: permissionRepository,
	}
}

type permissionService struct {
	*Service
	permissionRepository repository.PermissionRepository
}

func (s *permissionService) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	return s.permissionRepository.GetPermission(ctx, id)
}
