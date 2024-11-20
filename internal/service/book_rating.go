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
	GetBookRating(ctx context.Context, id uint) (*v1.BookRatingResponse, error)
	ListBookRatings(ctx context.Context, bookId uint, page, pageSize int) (*v1.ListBookRatingsResponse, error)
}

type bookRatingService struct {
	bookRatingRepo repository.BookRatingRepository
	*Service
}

func NewBookRatingService(service *Service, bookRatingRepo repository.BookRatingRepository) BookRatingService {
	return &bookRatingService{
		Service:        service,
		bookRatingRepo: bookRatingRepo,
	}
}

func (s *bookRatingService) CreateBookRating(ctx context.Context, req *v1.CreateBookRatingRequest) error {
	rating := &model.BookRating{
		BookId:       req.BookId,
		RatingTypeId: req.RatingTypeId,
		Comment:      req.Comment,
		IP:           req.IP,
	}
	return s.bookRatingRepo.Create(ctx, rating)
}

func (s *bookRatingService) UpdateBookRating(ctx context.Context, id uint, req *v1.UpdateBookRatingRequest) error {
	rating, err := s.bookRatingRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	rating.RatingTypeId = req.RatingTypeID
	rating.Comment = req.Comment

	return s.bookRatingRepo.Update(ctx, rating)
}

func (s *bookRatingService) DeleteBookRating(ctx context.Context, id uint) error {
	return s.bookRatingRepo.Delete(ctx, id)
}

func (s *bookRatingService) GetBookRating(ctx context.Context, id uint) (*v1.BookRatingResponse, error) {
	rating, err := s.bookRatingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1.BookRatingResponse{
		Id:           rating.Id,
		BookId:       rating.BookId,
		RatingTypeId: rating.RatingTypeId,
		Comment:      rating.Comment,
		IP:           rating.IP,
		CreatedAt:    rating.CreatedAt,
		UpdatedAt:    rating.UpdatedAt,
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
			RatingTypeId: rating.RatingTypeId,
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
