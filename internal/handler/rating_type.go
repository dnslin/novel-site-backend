package handler

import (
	"net/http"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingTypeHandler struct {
	*Handler
	ratingTypeService service.RatingTypeService
}

func NewRatingTypeHandler(handler *Handler, ratingTypeService service.RatingTypeService) *RatingTypeHandler {
	return &RatingTypeHandler{
		Handler:           handler,
		ratingTypeService: ratingTypeService,
	}
}

// CreateRatingType godoc
// @Summary 创建评分类型
// @Tags 评分类型模块
// @Accept json
// @Produce json
// @Param request body v1.CreateRatingTypeRequest true "params"
// @Success 200 {object} v1.Response
// @Router /rating-types [post]
// func (h *RatingTypeHandler) CreateRatingType(ctx *gin.Context) {
// 	req := new(v1.CreateRatingTypeRequest)
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
// 		return
// 	}

// 	if err := h.ratingTypeService.CreateRatingType(ctx, req); err != nil {
// 		h.logger.WithContext(ctx).Error("ratingTypeService.CreateRatingType error", zap.Error(err))
// 		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	v1.HandleSuccess(ctx, nil)
// }

// UpdateRatingType godoc
// @Summary 更新评分类型
// @Tags 评分类型模块
// @Accept json
// @Produce json
// @Param id path int true "评分类型ID"
// @Param request body v1.UpdateRatingTypeRequest true "params"
// @Success 200 {object} v1.Response
// @Router /rating-types/{id} [put]
// func (h *RatingTypeHandler) UpdateRatingType(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
// 		return
// 	}

// 	req := new(v1.UpdateRatingTypeRequest)
// 	if err := ctx.ShouldBindJSON(req); err != nil {
// 		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
// 		return
// 	}

// 	if err := h.ratingTypeService.UpdateRatingType(ctx, uint(id), req); err != nil {
// 		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	v1.HandleSuccess(ctx, nil)
// }

// DeleteRatingType godoc
// @Summary 删除评分类型
// @Tags 评分类型模块
// @Accept json
// @Produce json
// @Param id path int true "评分类型ID"
// @Success 200 {object} v1.Response
// @Router /rating-types/{id} [delete]
// func (h *RatingTypeHandler) DeleteRatingType(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
// 		return
// 	}

// 	if err := h.ratingTypeService.DeleteRatingType(ctx, uint(id)); err != nil {
// 		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	v1.HandleSuccess(ctx, nil)
// }

// GetRatingType godoc
// @Summary 获取评分类型详情
// @Tags 评分类型模块
// @Accept json
// @Produce json
// @Param id path int true "评分类型ID"
// @Success 200 {object} v1.RatingTypeResponse
// @Router /rating-types/{id} [get]
// func (h *RatingTypeHandler) GetRatingType(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
// 		return
// 	}

// 	ratingType, err := h.ratingTypeService.GetRatingType(ctx, uint(id))
// 	if err != nil {
// 		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
// 		return
// 	}

// 	v1.HandleSuccess(ctx, ratingType)
// }

// ListRatingTypes godoc
// @Summary 获取评分类型列表
// @Tags 评分类型模块
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} v1.ListRatingTypesResponse
// @Router /rating-types [get]
func (h *RatingTypeHandler) ListRatingTypes(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	ratingTypes, err := h.ratingTypeService.ListRatingTypes(ctx, page, pageSize)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, ratingTypes)
}
