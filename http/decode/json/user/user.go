package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"PRACTICESTUFF/example-go/domain"
	userEndpoint "PRACTICESTUFF/example-go/endpoints/user"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}
	return userEndpoint.FindRequest{UserID: userID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {

	nameQueryString := r.URL.Query().Get("name")

	return userEndpoint.FindAllRequest{
		Name: nameQueryString,
	}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}

	var req userEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.User.ID = userID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}
	return userEndpoint.DeleteRequest{UserID: userID}, nil
}