package service

import (
	"context"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type RatingTypeService interface {
	CreateRatingType(ctx context.Context, req *v1.CreateRatingTypeRequest) error
	UpdateRatingType(ctx context.Context, id uint, req *v1.UpdateRatingTypeRequest) error
	DeleteRatingType(ctx context.Context, id uint) error
	GetRatingType(ctx context.Context, id uint) (*v1.RatingTypeResponse, error)
	ListRatingTypes(ctx context.Context, page, pageSize int) (*v1.ListRatingTypesResponse, error)
}

type ratingTypeService struct {
	ratingTypeRepo repository.RatingTypeRepository
	*Service
}

func NewRatingTypeService(service *Service, ratingTypeRepo repository.RatingTypeRepository) RatingTypeService {
	return &ratingTypeService{
		Service:        service,
		ratingTypeRepo: ratingTypeRepo,
	}
}

// CreateRatingType 创建评分类型
func (s *ratingTypeService) CreateRatingType(ctx context.Context, req *v1.CreateRatingTypeRequest) error {
	return s.ratingTypeRepo.Create(ctx, &model.RatingType{
		Name:        req.Name,
		Description: req.Description,
	})
}

// UpdateRatingType 更新评分类型
func (s *ratingTypeService) UpdateRatingType(ctx context.Context, id uint, req *v1.UpdateRatingTypeRequest) error {
	rt, err := s.ratingTypeRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	rt.Name = req.Name
	rt.Description = req.Description
	rt.Level = req.Level

	return s.ratingTypeRepo.Update(ctx, rt)
}

// DeleteRatingType 删除评分类型
func (s *ratingTypeService) DeleteRatingType(ctx context.Context, id uint) error {
	return s.ratingTypeRepo.Delete(ctx, id)
}

// GetRatingType 获取评分类型详情
func (s *ratingTypeService) GetRatingType(ctx context.Context, id uint) (*v1.RatingTypeResponse, error) {
	ratingType, err := s.ratingTypeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &v1.RatingTypeResponse{
		Id:          ratingType.Id,
		Name:        ratingType.Name,
		Description: ratingType.Description,
		CreatedAt:   ratingType.CreatedAt,
		UpdatedAt:   ratingType.UpdatedAt,
	}, nil
}

// ListRatingTypes 获取评分类型列表
func (s *ratingTypeService) ListRatingTypes(ctx context.Context, page, pageSize int) (*v1.ListRatingTypesResponse, error) {
	ratingTypes, total, err := s.ratingTypeRepo.List(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	var items []*v1.RatingTypeResponse
	for _, rt := range ratingTypes {
		items = append(items, &v1.RatingTypeResponse{
			Id:          rt.Id,
			Name:        rt.Name,
			Description: rt.Description,
			CreatedAt:   rt.CreatedAt,
			UpdatedAt:   rt.UpdatedAt,
		})
	}

	return &v1.ListRatingTypesResponse{
		Total: total,
		Items: items,
	}, nil
}
