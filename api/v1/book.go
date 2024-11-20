package v1

import (
	"errors"
	"time"
)

var (
	ErrBookExists = errors.New("book already exists")
)

// Book相关
type CreateBookRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	FileSize    int64  `json:"file_size" binding:"required"`
	MD5         string `json:"md5" binding:"required"`
	NewFileName string `json:"new_file_name" binding:"required"`
	Cover       string `json:"cover"`
	Intro       string `json:"intro"`
	Parts       string `json:"parts"`
	FileURL     string `json:"file_url"`
	Sort        string `json:"sort"`
	Type        string `json:"type"`
	Tag         string `json:"tag"`
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Cover  string `json:"cover"`
	Intro  string `json:"intro"`
	Sort   string `json:"sort"`
	Type   string `json:"type"`
	Tag    string `json:"tag"`
}

type GetBookResponse struct {
	Id          uint      `json:"id"`
	FileName    string    `json:"file_name"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	FileSize    int64     `json:"file_size"`
	MD5         string    `json:"md5"`
	NewFileName string    `json:"new_file_name"`
	Cover       string    `json:"cover"`
	Intro       string    `json:"intro"`
	Parts       string    `json:"parts"`
	FileURL     string    `json:"file_url"`
	Sort        string    `json:"sort"`
	Type        string    `json:"type"`
	Tag         string    `json:"tag"`
	CreatedAt   time.Time `json:"created_at"`
}

type BookItem struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Cover     string    `json:"cover"`
	Intro     string    `json:"intro"`
	Sort      string    `json:"sort"`
	Type      string    `json:"type"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

type ListBooksResponse struct {
	Total int64       `json:"total"`
	Items []*BookItem `json:"items"`
}
