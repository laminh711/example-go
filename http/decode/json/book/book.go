package book

import (
	"context"
	"encoding/json"
	"net/http"

	"PRACTICESTUFF/example-go/domain"
	bookEndpoint "PRACTICESTUFF/example-go/endpoints/book"

	"github.com/go-chi/chi"
)

// FindRequest decode FindRequest
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.FindRequest{BookID: bookID}, nil
}

// FindAllRequest decode FindAllRequest
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return bookEndpoint.FindAllRequest{}, nil
}

// CreateRequest decode CreateRequest
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req bookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, err
}

// UpdateRequest decode UpdateRequest
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}

	var req bookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Book.ID = bookID

	return req, nil
}

// DeleteRequest decode DeleteRequest
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.DeleteRequest{ID: bookID}, nil
}
