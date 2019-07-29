package category

import (
	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// CreateData data for CreateCategory
type CreateData struct {
	Name string `json:"name"`
}

// CreateRequest request struct for CreateCategory
type CreateRequest struct {
	Category CreateData `json:"category"`
}

// CreateResponse response struct for CreateCategory
type CreateResponse struct {
	Category domain.Category `json:"category"`
}

// StatusCode override ok status code for success CreateCategory
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for CreateCategory
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			category = &domain.Category{
				Name: req.Category.Name,
			}
		)
		err := s.CategoryService.Create(ctx, category)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Category: *category}, nil
	}
}

// FindRequest request struct for FindCategory
type FindRequest struct {
	CategoryID domain.UUID
}

// FindResponse response struct for FindCategory
type FindResponse struct {
	Category *domain.Category `json:"category"`
}

// MakeFindEndpoint make endpoint for find Category
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       = request.(FindRequest)
			catToFind = domain.Category{Model: domain.Model{ID: req.CategoryID}}
		)

		category, err := s.CategoryService.Find(ctx, &catToFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Category: category}, nil
	}
}

// FindAllRequest request struct for FindAllCategory
type FindAllRequest struct{}

// FindAllResponse request struct for FindAllCategory
type FindAllResponse struct {
	Categories []domain.Category `json:"category"`
}

// MakeFindAllEndpoint make endpoint for FindAllCategory
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		categories, err := s.CategoryService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Categories: categories}, nil
	}
}

// UpdateData data for update
type UpdateData struct {
	ID   domain.UUID `json:"-"`
	Name string      `json:"name"`
}

// UpdateRequest request struct for UpdateCategory
type UpdateRequest struct {
	Category UpdateData `json:"category"`
}

// UpdateResponse response struct for UpdateCategory
type UpdateResponse struct {
	Category domain.Category `json:"category"`
}

// MakeUpdateEndpoint make endpoint for UpdateCategory
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		catToUpdate := domain.Category{
			Model: domain.Model{ID: req.Category.ID},
			Name:  req.Category.Name,
		}

		res, err := s.CategoryService.Update(ctx, &catToUpdate)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Category: *res}, nil
	}
}

// DeleteRequest request struct for DeteleCategory
type DeleteRequest struct {
	CategoryID domain.UUID
}

// DeleteResponse response struct for DeleteCategory
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for DeleteUser
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)

		catToDelete := domain.Category{
			Model: domain.Model{ID: req.CategoryID},
		}

		if err := s.CategoryService.Delete(ctx, &catToDelete); err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}
