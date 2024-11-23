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

// BookRatingStats 书籍评分统计
type BookRatingStats struct {
	BookId       uint                   `json:"book_id"`       // 书籍ID
	TotalRatings int64                  `json:"total_ratings"` // 总评分数
	RatingTypes  []*RatingTypeWithCount `json:"rating_types"`  // 各类型评分统计
}

// RatingTypeWithCount 评分类型及数量
type RatingTypeWithCount struct {
	Id          uint    `json:"id"`          // 评分类型ID
	Name        string  `json:"name"`        // 评分类型名称(仙草/粮草等)
	Description string  `json:"description"` // 评分类型描述
	Level       int     `json:"level"`       // 评分等级(5/4/3/2/1)
	Count       int64   `json:"count"`       // 该类型的评分数量
	Percentage  float64 `json:"percentage"`  // 该类型评分占比
}

// 修改 GetBookRating 的响应结构
type GetBookRatingResponse struct {
	Stats *BookRatingStats `json:"stats"`
}
