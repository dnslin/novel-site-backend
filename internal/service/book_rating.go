package service

import (
	"context"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type BookRatingService interface {
	CreateBookRating(ctx context.Context, req *v1.CreateBookRatingRequest) error
	UpdateBookRating(ctx context.Context, id uint, req *v1.UpdateBookRatingRequest) error
	DeleteBookRating(ctx context.Context, id uint) error
	GetBookRating(ctx context.Context, bookId uint) (*v1.GetBookRatingResponse, error)
	ListBookRatings(ctx context.Context, bookId uint, page, pageSize int) (*v1.ListBookRatingsResponse, error)
}

type bookRatingService struct {
	bookRatingRepo repository.BookRatingRepository
	ratingTypeRepo repository.RatingTypeRepository
	*Service
}

func NewBookRatingService(service *Service, bookRatingRepo repository.BookRatingRepository, ratingTypeRepo repository.RatingTypeRepository) BookRatingService {
	return &bookRatingService{
		Service:        service,
		bookRatingRepo: bookRatingRepo,
		ratingTypeRepo: ratingTypeRepo,
	}
}

func (s *bookRatingService) CreateBookRating(ctx context.Context, req *v1.CreateBookRatingRequest) error {
	rating := &model.BookRating{
		BookId:  req.BookId,
		Id:      req.RatingTypeId,
		Comment: req.Comment,
		IP:      req.IP,
	}
	return s.bookRatingRepo.Create(ctx, rating)
}

func (s *bookRatingService) UpdateBookRating(ctx context.Context, id uint, req *v1.UpdateBookRatingRequest) error {
	rating, err := s.bookRatingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	rating.Id = req.RatingTypeID
	rating.Comment = req.Comment

	return s.bookRatingRepo.Update(ctx, rating)
}

func (s *bookRatingService) DeleteBookRating(ctx context.Context, id uint) error {
	return s.bookRatingRepo.Delete(ctx, id)
}

func (s *bookRatingService) GetBookRating(ctx context.Context, bookId uint) (*v1.GetBookRatingResponse, error) {
	// 获取评分统计
	stats, total, err := s.bookRatingRepo.GetRatingStats(ctx, bookId)
	if err != nil {
		return nil, err
	}

	// 获取所有评分类型
	ratingTypes, _, err := s.ratingTypeRepo.List(ctx, 1, 100) // 假设评分类型不会超过100个
	if err != nil {
		return nil, err
	}

	// 构建评分类型统计map
	statsMap := make(map[uint]int64)
	for _, stat := range stats {
		statsMap[stat.RatingTypeID] = stat.Count
	}

	// 构建响应
	ratingTypeStats := make([]*v1.RatingTypeWithCount, 0)
	for _, rt := range ratingTypes {
		count := statsMap[rt.Id]
		percentage := float64(0)
		if total > 0 {
			percentage = float64(count) / float64(total) * 100
		}

		ratingTypeStats = append(ratingTypeStats, &v1.RatingTypeWithCount{
			Id:          rt.Id,
			Name:        rt.Name,
			Description: rt.Description,
			Level:       rt.Level,
			Count:       count,
			Percentage:  percentage,
		})
	}

	return &v1.GetBookRatingResponse{
		Stats: &v1.BookRatingStats{
			BookId:       bookId,
			TotalRatings: total,
			RatingTypes:  ratingTypeStats,
		},
	}, nil
}

func (s *bookRatingService) ListBookRatings(ctx context.Context, bookId uint, page, pageSize int) (*v1.ListBookRatingsResponse, error) {
	ratings, total, err := s.bookRatingRepo.ListByBookID(ctx, bookId, page, pageSize)
	if err != nil {
		return nil, err
	}

	var items []*v1.BookRatingResponse
	for _, rating := range ratings {
		items = append(items, &v1.BookRatingResponse{
			Id:           rating.Id,
			BookId:       rating.BookId,
			RatingTypeId: rating.Id,
			Comment:      rating.Comment,
			IP:           rating.IP,
			CreatedAt:    rating.CreatedAt,
			UpdatedAt:    rating.UpdatedAt,
		})
	}

	return &v1.ListBookRatingsResponse{
		Total: total,
		Items: items,
	}, nil
}
