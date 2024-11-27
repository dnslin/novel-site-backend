package handler

import (
	"github.com/gin-gonic/gin"
	"novel-site-backend/internal/service"
)

type RoleHandler struct {
	*Handler
	roleService service.RoleService
}

func NewRoleHandler(
    handler *Handler,
    roleService service.RoleService,
) *RoleHandler {
	return &RoleHandler{
		Handler:      handler,
		roleService: roleService,
	}
}

func (h *RoleHandler) GetRole(ctx *gin.Context) {

}
