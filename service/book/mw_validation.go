package book

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

func (mw validationMiddleware) Create(ctx context.Context, book *domain.Book) (err error) {
	if book.Name == "" {
		return ErrNameIsRequired
	}

	if len(book.Name) <= 5 {
		return ErrNameIsTooShort
	}

	if book.Description == "" {
		return ErrDescriptionIsRequired
	}

	if len(book.Description) <= 5 {
		return ErrDescriptionIsTooShort
	}

	var bookCat = domain.Category{Model: domain.Model{ID: book.CategoryID}}
	catExisted, err := mw.Service.IsCategoryExisted(ctx, &bookCat)
	if err != nil {
		return err
	}
	if !catExisted {
		return ErrCategoryNotExisted
	}

	return mw.Service.Create(ctx, book)
}

func (mw validationMiddleware) FindAll(ctx context.Context, queries FindAllQueries) ([]domain.Book, error) {
	return mw.Service.FindAll(ctx, queries)
}
func (mw validationMiddleware) Find(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	return mw.Service.Find(ctx, book)
}

func (mw validationMiddleware) Update(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book.Name == "" {
		return nil, ErrNameIsRequired
	}

	if len(book.Name) <= 5 {
		return nil, ErrNameIsTooShort
	}

	if book.Description == "" {
		return nil, ErrDescriptionIsRequired
	}

	if len(book.Description) <= 5 {
		return nil, ErrDescriptionIsTooShort
	}

	var bookCat = domain.Category{Model: domain.Model{ID: book.CategoryID}}
	catExisted, err := mw.Service.IsCategoryExisted(ctx, &bookCat)
	if err != nil {
		return nil, err
	}
	if !catExisted {
		return nil, ErrCategoryNotExisted
	}

	return mw.Service.Update(ctx, book)
}
func (mw validationMiddleware) Delete(ctx context.Context, book *domain.Book) error {
	return mw.Service.Delete(ctx, book)
}
