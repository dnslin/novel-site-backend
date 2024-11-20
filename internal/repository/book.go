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
	List(ctx context.Context, page, pageSize int) ([]*model.Book, int64, error)
	GetByMD5(ctx context.Context, md5 string) (*model.Book, error)
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

func (r *bookRepository) List(ctx context.Context, page, pageSize int) ([]*model.Book, int64, error) {
	var books []*model.Book
	var total int64

	db := r.DB(ctx)
	if err := db.Model(&model.Book{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
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
