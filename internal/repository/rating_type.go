package repository

import (
	"context"
	"novel-site-backend/internal/model"
)

type RatingTypeRepository interface {
	Create(ctx context.Context, rt *model.RatingType) error
	Update(ctx context.Context, rt *model.RatingType) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.RatingType, error)
	List(ctx context.Context, page, pageSize int) ([]*model.RatingType, int64, error)
}

type ratingTypeRepository struct {
	*Repository
}

func NewRatingTypeRepository(r *Repository) RatingTypeRepository {
	return &ratingTypeRepository{
		Repository: r,
	}
}

func (r *ratingTypeRepository) Create(ctx context.Context, rt *model.RatingType) error {
	if err := r.DB(ctx).Create(rt).Error; err != nil {
		return err
	}
	return nil
}

func (r *ratingTypeRepository) Update(ctx context.Context, rt *model.RatingType) error {
	if err := r.DB(ctx).Save(rt).Error; err != nil {
		return err
	}
	return nil
}

func (r *ratingTypeRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB(ctx).Delete(&model.RatingType{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ratingTypeRepository) GetByID(ctx context.Context, id uint) (*model.RatingType, error) {
	var rt model.RatingType
	if err := r.DB(ctx).First(&rt, id).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *ratingTypeRepository) List(ctx context.Context, page, pageSize int) ([]*model.RatingType, int64, error) {
	var rts []*model.RatingType
	var total int64

	offset := (page - 1) * pageSize

	if err := r.DB(ctx).Model(&model.RatingType{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB(ctx).Offset(offset).Limit(pageSize).Find(&rts).Error; err != nil {
		return nil, 0, err
	}

	return rts, total, nil
}
