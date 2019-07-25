package service

import (
	"PRACTICESTUFF/example-go/service/book"
	"PRACTICESTUFF/example-go/service/booklend"
	"PRACTICESTUFF/example-go/service/category"
	"PRACTICESTUFF/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	BooklendService booklend.Service
}
