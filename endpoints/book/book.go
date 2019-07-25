package book

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
)

// CreateData data for CreateBook
type CreateData struct {
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
	CategoryID  domain.UUID `json:"category_id"`
}

// CreateRequest request struct for CreateBook
type CreateRequest struct {
	Book CreateData `json:"book"`
}

// CreateResponse response struct for CreateBook
type CreateResponse struct {
	Book domain.Book `json:"book"`
}

// StatusCode customstatus code for success create Book
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Book
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		var (
			book = &domain.Book{
				Name:        req.Book.Name,
				Author:      req.Book.Author,
				Description: req.Book.Description,
				CategoryID:  req.Book.CategoryID,
			}
		)
		err := s.BookService.Create(ctx, book)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Book: *book}, nil
	}
}

// FindRequest request struct for finding a book
type FindRequest struct {
	BookID domain.UUID
}

// FindResponse response struct for finding a book
type FindResponse struct {
	Book domain.Book `json:"book"`
}

// MakeFindEndpoint make endpoint for finding a book
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindRequest)

		bookToFind := domain.Book{Model: domain.Model{ID: req.BookID}}

		res, err := s.BookService.Find(ctx, &bookToFind)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// FindAllRequest request struct for FindAllBook
type FindAllRequest struct{}

// FindAllResponse request struct for FindAllBook
type FindAllResponse struct {
	Books []domain.Book `json:"book"`
}

// MakeFindAllEndpoint make endpoint for FindAllBook
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		books, err := s.BookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Books: books}, nil
	}
}

// UpdateData data for update
type UpdateData struct {
	ID          domain.UUID `json:"-"`
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
	CategoryID  domain.UUID `json:"category_id"`
}

// UpdateRequest request struct for UpdateBook
type UpdateRequest struct {
	Book UpdateData `json:"book"`
}

// UpdateResponse response struct for UpdateBook
type UpdateResponse struct {
	Book domain.Book `json:"book"`
}

// MakeUpdateEndpoint make endpoint for UpdateBook
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		bookToUpdate := domain.Book{
			Model:       domain.Model{ID: req.Book.ID},
			Name:        req.Book.Name,
			Author:      req.Book.Author,
			Description: req.Book.Description,
			CategoryID:  req.Book.CategoryID,
		}

		res, err := s.BookService.Update(ctx, &bookToUpdate)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Book: *res}, nil
	}
}

// DeleteRequest request struct for DeteleBook
type DeleteRequest struct {
	ID domain.UUID
}

// DeleteResponse response struct for DeleteBook
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for DeleteUser
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		bookToDelete := domain.Book{Model: domain.Model{ID: req.ID}}
		if err := s.BookService.Delete(ctx, &bookToDelete); err != nil {
			return nil, err
		}
		return DeleteResponse{"success"}, nil
	}
}
