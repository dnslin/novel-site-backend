package v1

import "time"

type CreateBookRatingRequest struct {
	BookId       uint   `json:"book_id" binding:"required"`
	RatingTypeId uint   `json:"rating_type_id" binding:"required"`
	Comment      string `json:"comment"`
	IP           string `json:"ip"`
}

type UpdateBookRatingRequest struct {
	Comment      string `json:"comment"`
	RatingTypeID uint   `json:"rating_type_id"`
}

type BookRatingResponse struct {
	Id           uint      `json:"id"`
	BookId       uint      `json:"book_id"`
	RatingTypeId uint      `json:"rating_type_id"`
	Comment      string    `json:"comment"`
	IP           string    `json:"ip"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ListBookRatingsResponse struct {
	Total int64                 `json:"total"`
	Items []*BookRatingResponse `json:"items"`
}
