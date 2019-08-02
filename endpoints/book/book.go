package book

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
	bookService "PRACTICESTUFF/example-go/service/book"
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
	Book []CreateData `json:"book"`
}

// CreateResponse response struct for CreateBook
type CreateResponse struct {
	Book []domain.Book `json:"book"`
}

// StatusCode customstatus code for success create Book
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Book
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		res := CreateResponse{
			[]domain.Book{},
		}

		req := request.(CreateRequest)

		inpData := []domain.Book{}
		for _, cdata := range req.Book {
			book := &domain.Book{
				Name:        cdata.Name,
				Author:      cdata.Author,
				Description: cdata.Description,
				CategoryID:  cdata.CategoryID,
			}
			inpData = append(inpData, *book)
		}

		sth, err := s.BookService.CreateBatch(ctx, inpData)
		if err != nil {
			return nil, err
		}

		res.Book = inpData

		// for _, cdata := range req.Book {
		// 	book := &domain.Book{
		// 		Name:        cdata.Name,
		// 		Author:      cdata.Author,
		// 		Description: cdata.Description,
		// 		CategoryID:  cdata.CategoryID,
		// 	}

		// 	err := s.BookService.Create(ctx, book)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	res.Book = append(res.Book, *book)
		// }

		return CreateResponse{Book: sth}, nil
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
type FindAllRequest struct {
	Name    string
	Status  string
	TagName string
}

// FindAllResponse request struct for FindAllBook
type FindAllResponse struct {
	Books []domain.Book `json:"book"`
}

// MakeFindAllEndpoint make endpoint for FindAllBook
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindAllRequest)
		books, err := s.BookService.FindAll(ctx, bookService.FindAllQueries(req))
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
	BookID domain.UUID
}

// DeleteResponse response struct for DeleteBook
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for DeleteUser
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		bookToDelete := domain.Book{Model: domain.Model{ID: req.BookID}}
		if err := s.BookService.Delete(ctx, &bookToDelete); err != nil {
			return nil, err
		}
		return DeleteResponse{"success"}, nil
	}
}

// AddTagsData data for AddTags
type AddTagsData struct {
	TagID domain.UUID `json:"tag_id"`
}

// AddTagsRequest request struct for AddTags
type AddTagsRequest struct {
	BookID domain.UUID   `json:"-"`
	Tag    []AddTagsData `json:"tag"`
}

// AddTagsResponse response struct for AddTags
type AddTagsResponse struct {
	Status string `json:"status"`
}

// StatusCode customstatus code for success create Book
func (AddTagsResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeAddTagsToBookEndpoint make endpoint for create a Book
func MakeAddTagsToBookEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(AddTagsRequest)

		inpData := []domain.Tag{}
		for _, cdata := range req.Tag {
			tag := domain.Tag{
				Model: domain.Model{ID: cdata.TagID},
			}
			inpData = append(inpData, tag)
		}

		book := domain.Book{Model: domain.Model{ID: req.BookID}}

		_, err := s.BookService.AddTags(ctx, &book, inpData)
		if err != nil {
			return nil, err
		}

		return AddTagsResponse{Status: "success"}, nil
	}
}
