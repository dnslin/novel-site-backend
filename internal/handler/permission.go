package handler

import (
	"github.com/gin-gonic/gin"
	"novel-site-backend/internal/service"
)

type PermissionHandler struct {
	*Handler
	permissionService service.PermissionService
}

func NewPermissionHandler(
    handler *Handler,
    permissionService service.PermissionService,
) *PermissionHandler {
	return &PermissionHandler{
		Handler:      handler,
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) GetPermission(ctx *gin.Context) {

}
