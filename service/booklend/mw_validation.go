package booklend

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

func (mw validationMiddleware) Create(ctx context.Context, booklend *domain.Booklend) (err error) {
	user := domain.User{Model: domain.Model{ID: booklend.UserID}}
	userExisted, err := mw.Service.IsUserExisted(ctx, &user)
	if err != nil {
		return err
	}
	if !userExisted {
		return ErrUserNotFound
	}

	book := domain.Book{Model: domain.Model{ID: booklend.BookID}}
	bookExisted, err := mw.Service.IsBookExisted(ctx, &book)
	if err != nil {
		return err
	}
	if !bookExisted {
		return ErrBookNotFound
	}

	bookLendable, err := mw.Service.IsBooklendable(ctx, booklend)
	if err != nil {
		return err
	}
	if !bookLendable {
		return ErrBookNotLendable
	}

	if booklend.From.After(booklend.To) {
		return ErrInvalidTimeSpan
	}

	return mw.Service.Create(ctx, booklend)
}
