package handler

import (
	"github.com/gin-gonic/gin"
	"novel-site-backend/internal/service"
)

type UserRoleHandler struct {
	*Handler
	userRoleService service.UserRoleService
}

func NewUserRoleHandler(
    handler *Handler,
    userRoleService service.UserRoleService,
) *UserRoleHandler {
	return &UserRoleHandler{
		Handler:      handler,
		userRoleService: userRoleService,
	}
}

func (h *UserRoleHandler) GetUserRole(ctx *gin.Context) {

}
