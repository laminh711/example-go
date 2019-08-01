package book

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

// FindAllQueries query struct got from FindAllRequest
type FindAllQueries struct {
	Name   string
	Status string
}

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Book) error
	Find(ctx context.Context, p *domain.Book) (*domain.Book, error)
	FindAll(ctx context.Context, queries FindAllQueries) ([]domain.Book, error)
	Update(ctx context.Context, p *domain.Book) (*domain.Book, error)
	Delete(ctx context.Context, p *domain.Book) error
	IsCategoryExisted(ctx context.Context, cat *domain.Category) (bool, error)
	CreateBatch(ctx context.Context, p []domain.Book) error
	AddTags(ctx context.Context, p *domain.Book, t []domain.Tag) ([]domain.BookTag, error)
}
