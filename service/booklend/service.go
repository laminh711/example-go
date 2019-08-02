package booklend

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Booklend) error
	IsUserExisted(ctx context.Context, p *domain.User) (bool, error)
	IsBookExisted(ctx context.Context, p *domain.Book) (bool, error)
	IsBooklendable(ctx context.Context, p *domain.Booklend) (bool, error)

	CreateBatch(ctx context.Context, p []domain.Booklend) ([]domain.Booklend, error)
}
