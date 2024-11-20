package repository

import (
	"context"
	"novel-site-backend/internal/model"
)

type BookRatingRepository interface {
	Create(ctx context.Context, br *model.BookRating) error
	Update(ctx context.Context, br *model.BookRating) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.BookRating, error)
	ListByBookID(ctx context.Context, bookId uint, page, pageSize int) ([]*model.BookRating, int64, error)
}

type bookRatingRepository struct {
	*Repository
}

func NewBookRatingRepository(r *Repository) BookRatingRepository {
	return &bookRatingRepository{
		Repository: r,
	}
}

func (r *bookRatingRepository) Create(ctx context.Context, br *model.BookRating) error {
	if err := r.DB(ctx).Create(br).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRatingRepository) Update(ctx context.Context, br *model.BookRating) error {
	if err := r.DB(ctx).Save(br).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRatingRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB(ctx).Delete(&model.BookRating{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRatingRepository) GetByID(ctx context.Context, id uint) (*model.BookRating, error) {
	var br model.BookRating
	if err := r.DB(ctx).First(&br, id).Error; err != nil {
		return nil, err
	}
	return &br, nil
}

func (r *bookRatingRepository) ListByBookID(ctx context.Context, bookId uint, page, pageSize int) ([]*model.BookRating, int64, error) {
	var ratings []*model.BookRating
	var total int64

	offset := (page - 1) * pageSize

	if err := r.DB(ctx).Model(&model.BookRating{}).Where("book_id = ?", bookId).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB(ctx).Where("book_id = ?", bookId).Offset(offset).Limit(pageSize).Find(&ratings).Error; err != nil {
		return nil, 0, err
	}

	return ratings, total, nil
}
