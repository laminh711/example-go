package booklend

import "net/http"

// Error declaration
var (
	ErrNotFound         = errNotFound{}
	ErrUnknown          = errUnknown{}
	ErrBookNotFound     = errBookNotFound{}
	ErrBookNotAvailable = errBookNotAvailable{}
	ErrUserNotFound     = errUserNotFound{}
	ErrInvalidTimeSpan  = errInvalidTimeSpan{}
	ErrRecordNotFound   = errRecordNotFound{}
)

type errInvalidTimeSpan struct{}

func (errInvalidTimeSpan) Error() string {
	return `'from' must be sooner than 'to'`
}

func (errInvalidTimeSpan) StatusCode() int {
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

type errUserNotFound struct{}

func (errUserNotFound) Error() string {
	return "invalid user id"
}

func (errUserNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNotFound struct{}

func (errBookNotFound) Error() string {
	return "invalid book id"
}

func (errBookNotFound) StatusCode() int {
	return http.StatusBadRequest
}

type errBookNotAvailable struct{}

func (errBookNotAvailable) Error() string {
	return "book not available"
}

func (errBookNotAvailable) StatusCode() int {
	return http.StatusConflict
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}
