package service

import (
	"context"
	v1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/model"
	"novel-site-backend/internal/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, req *v1.CreateBookRequest) error
	UpdateBook(ctx context.Context, id uint, req *v1.UpdateBookRequest) error
	DeleteBook(ctx context.Context, id uint) error
	GetBook(ctx context.Context, id uint) (*v1.GetBookResponse, error)
	ListBooks(ctx context.Context, page, pageSize int) (*v1.ListBooksResponse, error)
}

type bookService struct {
	bookRepo repository.BookRepository
	*Service
}

func NewBookService(service *Service, bookRepo repository.BookRepository) BookService {
	return &bookService{
		Service:  service,
		bookRepo: bookRepo,
	}
}

func (s *bookService) CreateBook(ctx context.Context, req *v1.CreateBookRequest) error {
	// 检查MD5是否已存在
	existBook, err := s.bookRepo.GetByMD5(ctx, req.MD5)
	if err != nil {
		return err
	}
	if existBook != nil {
		return v1.ErrBookExists
	}

	book := &model.Book{
		FileName:    req.FileName,
		Title:       req.Title,
		Author:      req.Author,
		FileSize:    req.FileSize,
		MD5:         req.MD5,
		NewFileName: req.NewFileName,
		Cover:       req.Cover,
		Intro:       req.Intro,
		Parts:       req.Parts,
		FileURL:     req.FileURL,
		Sort:        req.Sort,
		Type:        req.Type,
		Tag:         req.Tag,
	}

	return s.bookRepo.Create(ctx, book)
}

func (s *bookService) UpdateBook(ctx context.Context, id uint, req *v1.UpdateBookRequest) error {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	book.Title = req.Title
	book.Author = req.Author
	book.Cover = req.Cover
	book.Intro = req.Intro
	book.Sort = req.Sort
	book.Type = req.Type
	book.Tag = req.Tag

	return s.bookRepo.Update(ctx, book)
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {
	return s.bookRepo.Delete(ctx, id)
}

func (s *bookService) GetBook(ctx context.Context, id uint) (*v1.GetBookResponse, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1.GetBookResponse{
		Id:          book.Id,
		FileName:    book.FileName,
		Title:       book.Title,
		Author:      book.Author,
		FileSize:    book.FileSize,
		MD5:         book.MD5,
		NewFileName: book.NewFileName,
		Cover:       book.Cover,
		Intro:       book.Intro,
		Parts:       book.Parts,
		FileURL:     book.FileURL,
		Sort:        book.Sort,
		Type:        book.Type,
		Tag:         book.Tag,
		CreatedAt:   book.CreatedAt,
	}, nil
}

func (s *bookService) ListBooks(ctx context.Context, page, pageSize int) (*v1.ListBooksResponse, error) {
	books, total, err := s.bookRepo.List(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	var items []*v1.BookItem
	for _, book := range books {
		items = append(items, &v1.BookItem{
			Id:        book.Id,
			Title:     book.Title,
			Author:    book.Author,
			Cover:     book.Cover,
			Intro:     book.Intro,
			Sort:      book.Sort,
			Type:      book.Type,
			Tag:       book.Tag,
			CreatedAt: book.CreatedAt,
		})
	}

	return &v1.ListBooksResponse{
		Total: total,
		Items: items,
	}, nil
}