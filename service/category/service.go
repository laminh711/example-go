package category

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Category) error
	Update(ctx context.Context, p *domain.Category) (*domain.Category, error)
	Find(ctx context.Context, p *domain.Category) (*domain.Category, error)
	FindAll(ctx context.Context, nameToFind string) ([]domain.Category, error)
	Delete(ctx context.Context, p *domain.Category) error

	IsNameDuplicate(ctx context.Context, nameToSearch string) (bool, error)
}
