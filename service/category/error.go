package category

import "net/http"

// Error declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrNameIsRequired   = errNameIsRequired{}
	ErrNameIsTooShort   = errNameIsTooShort{}
	ErrNameIsDuplicated = errNameIsDuplicated{}
	ErrRecordNotFound   = errRecordNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}

func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}

func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "category name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsTooShort struct{}

func (errNameIsTooShort) Error() string {
	return "category name must be at least 5 characters long"
}

func (errNameIsTooShort) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsDuplicated struct{}

func (errNameIsDuplicated) Error() string {
	return "category name is duplicated"
}

func (errNameIsDuplicated) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}
