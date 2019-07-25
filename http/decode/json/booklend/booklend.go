package booklend

import (
	booklendEndpoint "PRACTICESTUFF/example-go/endpoints/booklend"
	"context"
	"encoding/json"
	"net/http"
)

// CreateRequest decode create a BookLend request
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req booklendEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, err
}
