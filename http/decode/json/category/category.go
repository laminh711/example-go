package category

import (
	"PRACTICESTUFF/example-go/domain"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	categoryEndpoint "PRACTICESTUFF/example-go/endpoints/category"
)

// FindRequest decode FindRequest
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.FindRequest{CategoryID: categoryID}, nil
}

// FindAllRequest decode FindAllRequest
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	nameToFind := r.URL.Query().Get("name")
	return categoryEndpoint.FindAllRequest{Name: nameToFind}, nil
}

// CreateRequest decode CreateRequest
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req categoryEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest decode UpdateRequest
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}

	var req categoryEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Category.ID = categoryID

	return req, nil
}

// DeleteRequest decode DeleteRequest
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))

	if err != nil {
		return nil, err
	}

	return categoryEndpoint.DeleteRequest{ID: categoryID}, nil
}
