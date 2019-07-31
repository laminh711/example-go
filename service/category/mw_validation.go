package category

import (
	"context"

	"PRACTICESTUFF/example-go/domain"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, category *domain.Category) (err error) {
	if category.Name == "" {
		return ErrNameIsRequired
	}

	if len(category.Name) <= 5 {
		return ErrNameIsTooShort
	}

	dup, err := mw.Service.IsNameDuplicate(ctx, category.Name)
	if err != nil {
		return err
	}
	if dup {
		return ErrNameIsDuplicated
	}

	return mw.Service.Create(ctx, category)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Category, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return mw.Service.Find(ctx, category)
}

func (mw validationMiddleware) Update(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if category.Name == "" {
		return nil, ErrNameIsRequired
	}

	if len(category.Name) <= 5 {
		return nil, ErrNameIsTooShort
	}

	dup, err := mw.Service.IsNameDuplicate(ctx, category.Name)
	if err != nil {
		return nil, err
	}
	if dup {
		return nil, ErrNameIsDuplicated
	}

	return mw.Service.Update(ctx, category)
}
func (mw validationMiddleware) Delete(ctx context.Context, category *domain.Category) error {
	return mw.Service.Delete(ctx, category)
}
