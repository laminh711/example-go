package book

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Book) error
	Find(ctx context.Context, p *domain.Book) (*domain.Book, error)
	FindAll(ctx context.Context) ([]domain.Book, error)
	Update(ctx context.Context, p *domain.Book) (*domain.Book, error)
	Delete(ctx context.Context, p *domain.Book) error

	IsCategoryExisted(ctx context.Context, cat *domain.Category) (bool, error)
}
