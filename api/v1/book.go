package v1

import (
	"errors"
	"time"
)

var (
	ErrBookExists = errors.New("book already exists")
)

// CreateBookRequest 创建图书请求
type CreateBookRequest struct {
	FileName    string `json:"file_name" binding:"required"`     // 原始文件名
	Title       string `json:"title" binding:"required"`         // 书名
	Author      string `json:"author" binding:"required"`        // 作者
	FileSize    int64  `json:"file_size" binding:"required"`     // 文件大小（字节）
	MD5         string `json:"md5" binding:"required"`           // 文件MD5值
	NewFileName string `json:"new_file_name" binding:"required"` // 新文件名
	Cover       string `json:"cover"`                            // 封面图片URL
	Intro       string `json:"intro"`                            // 简介
	Parts       string `json:"parts"`                            // 章节信息
	FileURL     string `json:"file_url"`                         // 文件URL
	Sort        string `json:"sort"`                             // 分类
	Type        string `json:"type"`                             // 类型
	Tag         string `json:"tag"`                              // 标签
}

// UpdateBookRequest 更新图书请求
type UpdateBookRequest struct {
	Title  string `json:"title"`  // 书名
	Author string `json:"author"` // 作者
	Cover  string `json:"cover"`  // 封面图片URL
	Intro  string `json:"intro"`  // 简介
	Sort   string `json:"sort"`   // 分类
	Type   string `json:"type"`   // 类型
	Tag    string `json:"tag"`    // 标签
}

// GetBookResponse 获取图书响应
type GetBookResponse struct {
	Id          uint      `json:"id"`            // 图书ID
	FileName    string    `json:"file_name"`     // 原始文件名
	Title       string    `json:"title"`         // 书名
	Author      string    `json:"author"`        // 作者
	FileSize    int64     `json:"file_size"`     // 文件大小（字节）
	MD5         string    `json:"md5"`           // 文件MD5值
	NewFileName string    `json:"new_file_name"` // 新文件名
	Cover       string    `json:"cover"`         // 封面图片URL
	Intro       string    `json:"intro"`         // 简介
	Parts       string    `json:"parts"`         // 章节信息
	FileURL     string    `json:"file_url"`      // 文件URL
	Sort        string    `json:"sort"`          // 分类
	Type        string    `json:"type"`          // 类型
	Tag         string    `json:"tag"`           // 标签
	HotValue    int64     `json:"hot_value"`     // 热度值
	CreatedAt   time.Time `json:"created_at"`    // 创建时间
}

// BookItem 图书列表项
type BookItem struct {
	Id        uint      `json:"id"`         // 图书ID
	Title     string    `json:"title"`      // 书名
	Author    string    `json:"author"`     // 作者
	Cover     string    `json:"cover"`      // 封面图片URL
	Intro     string    `json:"intro"`      // 简介
	Sort      string    `json:"sort"`       // 分类
	Type      string    `json:"type"`       // 类型
	Tag       string    `json:"tag"`        // 标签
	HotValue  int64     `json:"hot_value"`  // 热度值
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

type ListBooksRequest struct {
	Title    string `json:"title,omitempty"`  // 书名，可选，支持模糊查询
	Author   string `json:"author,omitempty"` // 作者，可选，支持模糊查询
	Tag      string `json:"tag,omitempty"`    // 标签，可选，支持模糊查询
	Page     int    `json:"page"`             // 页码
	PageSize int    `json:"page_size"`        // 每页数量
}

type ListBooksResponse struct {
	Total int64       `json:"total"`
	Items []*BookItem `json:"items"`
}
