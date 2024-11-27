package handler

import (
	"net/http"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register 处理用户注册请求
// 1. 验证请求参数
// 2. 调用service层注册用户
// 3. 返回结果
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.logger.WithContext(ctx).Warn("注册参数无效", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("用户注册失败",
			zap.String("邮箱", req.Email),
			zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.logger.WithContext(ctx).Info("用户注册成功",
		zap.String("邮箱", req.Email))
	v1.HandleSuccess(ctx, nil)
}

// Login 处理用户登录请求
// 1. 验证登录参数
// 2. 调用service层验证用户身份并生成token
// 3. 返回token
func (h *UserHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.WithContext(ctx).Warn("登录参数无效", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		h.logger.WithContext(ctx).Error("用户登录失败",
			zap.String("邮箱", req.Email),
			zap.Error(err))
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	h.logger.WithContext(ctx).Info("用户登录成功",
		zap.String("邮箱", req.Email))
	v1.HandleSuccess(ctx, v1.LoginResponseData{AccessToken: token})
}

// GetProfile 获取用户信息
// 1. 从上下文获取用户ID
// 2. 调用service层获取用户信息
// 3. 返回用户信息
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		h.logger.WithContext(ctx).Warn("未授权的个人信息访问")
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		h.logger.WithContext(ctx).Error("获取用户信息失败",
			zap.String("用户ID", userId),
			zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	h.logger.WithContext(ctx).Info("获取用户信息成功",
		zap.String("用户ID", userId))
	v1.HandleSuccess(ctx, user)
}

// UpdateProfile 更新用户信息
// 1. 从上下文获取用户ID
// 2. 验证更新参数
// 3. 调用service层更新用户信息
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		h.logger.WithContext(ctx).Warn("未授权的信息更新访问")
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	var req v1.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.WithContext(ctx).Warn("更新参数无效", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		h.logger.WithContext(ctx).Error("更新用户信息失败",
			zap.String("用户ID", userId),
			zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	h.logger.WithContext(ctx).Info("更新用户信息成功",
		zap.String("用户ID", userId))
	v1.HandleSuccess(ctx, nil)
}
