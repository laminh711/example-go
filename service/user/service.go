package user

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.User) error
	Update(ctx context.Context, p *domain.User) (*domain.User, error)
	Find(ctx context.Context, p *domain.User) (*domain.User, error)
	FindAll(ctx context.Context, nameToFind string) ([]domain.User, error)
	Delete(ctx context.Context, p *domain.User) error
}
