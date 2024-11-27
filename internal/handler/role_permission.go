package handler

import (
	"github.com/gin-gonic/gin"
	"novel-site-backend/internal/service"
)

type RolePermissionHandler struct {
	*Handler
	rolePermissionService service.RolePermissionService
}

func NewRolePermissionHandler(
    handler *Handler,
    rolePermissionService service.RolePermissionService,
) *RolePermissionHandler {
	return &RolePermissionHandler{
		Handler:      handler,
		rolePermissionService: rolePermissionService,
	}
}

func (h *RolePermissionHandler) GetRolePermission(ctx *gin.Context) {

}
