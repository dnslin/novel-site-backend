package handler

import (
	"net/http"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookHandler struct {
	*Handler
	bookService service.BookService
}

func NewBookHandler(handler *Handler, bookService service.BookService) *BookHandler {
	return &BookHandler{
		Handler:     handler,
		bookService: bookService,
	}
}

// CreateBook godoc
// @Summary 创建书籍
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param request body v1.CreateBookRequest true "params"
// @Success 200 {object} v1.Response
// @Router /books [post]
func (h *BookHandler) CreateBook(ctx *gin.Context) {
	req := new(v1.CreateBookRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.bookService.CreateBook(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("bookService.CreateBook error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// UpdateBook godoc
// @Summary 更新书籍
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param id path int true "书籍ID"
// @Param request body v1.UpdateBookRequest true "params"
// @Success 200 {object} v1.Response
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	req := new(v1.UpdateBookRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.bookService.UpdateBook(ctx, uint(id), req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteBook godoc
// @Summary 删除书籍
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param id path int true "书籍ID"
// @Success 200 {object} v1.Response
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.bookService.DeleteBook(ctx, uint(id)); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// GetBook godoc
// @Summary 获取书籍详情
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param id path int true "书籍ID"
// @Success 200 {object} v1.GetBookResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	book, err := h.bookService.GetBook(ctx, uint(id))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, book)
}

// ListBooks godoc
// @Summary 获取书籍列表
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param request body v1.ListBooksRequest true "查询参数"
// @Success 200 {object} v1.ListBooksResponse
// @Router /books/list [post]
func (h *BookHandler) ListBooks(ctx *gin.Context) {
	req := new(v1.ListBooksRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	books, err := h.bookService.ListBooks(ctx, req)
	if err != nil {
		h.logger.WithContext(ctx).Error("bookService.ListBooks error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, books)
}

// GetAllSorts godoc
// @Summary 获取所有书籍分类
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Success 200 {object} v1.GetAllSortsResponse
// @Router /books/sorts [get]
func (h *BookHandler) GetAllSorts(ctx *gin.Context) {
	sorts, err := h.bookService.GetAllSorts(ctx)
	if err != nil {
		h.logger.WithContext(ctx).Error("bookService.GetAllSorts error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, &v1.GetAllSortsResponse{
		Sorts: sorts,
	})
}

// QuickSearch godoc
// @Summary 快速搜索书籍
// @Tags 书籍模块
// @Accept json
// @Produce json
// @Param request body v1.QuickSearchRequest true "搜索参数"
// @Success 200 {object} v1.QuickSearchResponse
// @Router /books/search [post]
func (h *BookHandler) QuickSearch(ctx *gin.Context) {
	req := new(v1.QuickSearchRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	result, err := h.bookService.QuickSearch(ctx, req.Keyword)
	if err != nil {
		h.logger.WithContext(ctx).Error("bookService.QuickSearch error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, result)
}
