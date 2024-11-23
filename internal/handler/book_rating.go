package handler

import (
	"net/http"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/service"
	"strconv"

	"novel-site-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookRatingHandler struct {
	*Handler
	bookRatingService service.BookRatingService
}

func NewBookRatingHandler(handler *Handler, bookRatingService service.BookRatingService) *BookRatingHandler {
	return &BookRatingHandler{
		Handler:           handler,
		bookRatingService: bookRatingService,
	}
}

// CreateBookRating godoc
// @Summary 创建书籍评分
// @Tags 书籍评分模块
// @Accept json
// @Produce json
// @Param request body v1.CreateBookRatingRequest true "params"
// @Success 200 {object} v1.Response
// @Router /book-ratings [post]
func (h *BookRatingHandler) CreateBookRating(ctx *gin.Context) {
	req := new(v1.CreateBookRatingRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 从中间件获取IP
	req.IP = middleware.GetClientIP(ctx)

	if err := h.bookRatingService.CreateBookRating(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("bookRatingService.CreateBookRating error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// UpdateBookRating godoc
// @Summary 更新书籍评分
// @Tags 书籍评分模块
// @Accept json
// @Produce json
// @Param id path int true "评分ID"
// @Param request body v1.UpdateBookRatingRequest true "params"
// @Success 200 {object} v1.Response
// @Router /book-ratings/{id} [put]
func (h *BookRatingHandler) UpdateBookRating(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	req := new(v1.UpdateBookRatingRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.bookRatingService.UpdateBookRating(ctx, uint(id), req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteBookRating godoc
// @Summary 删除书籍评分
// @Tags 书籍评分模块
// @Accept json
// @Produce json
// @Param id path int true "评分ID"
// @Success 200 {object} v1.Response
// @Router /book-ratings/{id} [delete]
func (h *BookRatingHandler) DeleteBookRating(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.bookRatingService.DeleteBookRating(ctx, uint(id)); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// GetBookRating godoc
// @Summary 获取书籍评分统计
// @Tags 书籍评分模块
// @Accept json
// @Produce json
// @Param book_id path int true "书籍ID"
// @Success 200 {object} v1.GetBookRatingResponse
// @Router /books/{book_id}/rating-stats [get]
func (h *BookRatingHandler) GetBookRating(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("book_id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	stats, err := h.bookRatingService.GetBookRating(ctx, uint(bookId))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, stats)
}

// ListBookRatings godoc
// @Summary 获取书籍评分列表
// @Tags 书籍评分模块
// @Accept json
// @Produce json
// @Param book_id path int true "书籍ID"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} v1.ListBookRatingsResponse
// @Router /books/{book_id}/ratings [get]
func (h *BookRatingHandler) ListBookRatings(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("book_id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	ratings, err := h.bookRatingService.ListBookRatings(ctx, uint(bookId), page, pageSize)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, ratings)
}
