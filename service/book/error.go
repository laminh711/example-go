package book

import "net/http"

// Error declaration
var (
	ErrNotFound              = errNotFound{}
	ErrUnknown               = errUnknown{}
	ErrNameIsRequired        = errNameIsRequired{}
	ErrNameIsTooShort        = errNameIsTooShort{}
	ErrDescriptionIsRequired = errDescriptionIsRequired{}
	ErrDescriptionIsTooShort = errDescriptionIsTooShort{}
	ErrRecordNotFound        = errRecordNotFound{}
	ErrCategoryNotExisted    = errCategoryNotExisted{}
)

type errDescriptionIsRequired struct{}

func (errDescriptionIsRequired) Error() string {
	return "book description is required"
}

func (errDescriptionIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errDescriptionIsTooShort struct{}

func (errDescriptionIsTooShort) Error() string {
	return "book description must be at least 5 characters long"
}

func (errDescriptionIsTooShort) StatusCode() int {
	return http.StatusBadRequest
}

type errCategoryNotExisted struct{}

func (errCategoryNotExisted) Error() string {
	return "invalid category id"
}

func (errCategoryNotExisted) StatusCode() int {
	return http.StatusBadRequest
}

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
	return "book name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsTooShort struct{}

func (errNameIsTooShort) Error() string {
	return "book name must be at least 5 characters long"
}

func (errNameIsTooShort) StatusCode() int {
	return http.StatusBadRequest
}

type errNameIsDuplicated struct{}

func (errNameIsDuplicated) Error() string {
	return "book name is duplicated"
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
