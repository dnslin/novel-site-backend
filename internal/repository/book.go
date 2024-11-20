package repository

import (
	"context"
	"errors"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Book, error)
	List(ctx context.Context, req *v1.ListBooksRequest) ([]*model.Book, int64, error)
	GetByMD5(ctx context.Context, md5 string) (*model.Book, error)
	IncrementHotValue(ctx context.Context, id uint) error
}

type bookRepository struct {
	*Repository
}

func NewBookRepository(r *Repository) BookRepository {
	return &bookRepository{
		Repository: r,
	}
}

func (r *bookRepository) Create(ctx context.Context, book *model.Book) error {
	return r.DB(ctx).Create(book).Error
}

func (r *bookRepository) Update(ctx context.Context, book *model.Book) error {
	return r.DB(ctx).Save(book).Error
}

func (r *bookRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&model.Book{}, id).Error
}

func (r *bookRepository) GetByID(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := r.DB(ctx).First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) List(ctx context.Context, req *v1.ListBooksRequest) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64

	db := r.DB(ctx)

	// 构建查询条件
	query := db.Model(&model.Book{})

	// 添加模糊查询条件
	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.Author != "" {
		query = query.Where("author LIKE ?", "%"+req.Author+"%")
	}
	if req.Tag != "" {
		query = query.Where("tag LIKE ?", "%"+req.Tag+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询，按热度值降序排序
	offset := (req.Page - 1) * req.PageSize
	if err := query.Order("hot_value DESC").
		Offset(offset).
		Limit(req.PageSize).
		Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (r *bookRepository) GetByMD5(ctx context.Context, md5 string) (*model.Book, error) {
	var book model.Book
	if err := r.DB(ctx).Where("md5 = ?", md5).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) IncrementHotValue(ctx context.Context, id uint) error {
	// 使用 SQL 的 UPDATE 语句直接增加热度值
	return r.DB(ctx).Model(&model.Book{}).
		Where("id = ?", id).
		UpdateColumn("hot_value", gorm.Expr("hot_value + ?", 1)).
		Error
}
