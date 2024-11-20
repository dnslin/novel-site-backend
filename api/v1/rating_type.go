package v1

import "time"

type CreateRatingTypeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Level       int    `json:"level" binding:"required"`
}

type UpdateRatingTypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type RatingTypeResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Level       int       `json:"level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListRatingTypesResponse struct {
	Total int64                 `json:"total"`
	Items []*RatingTypeResponse `json:"items"`
}
