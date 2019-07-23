package service

import (
	"PRACTICESTUFF/example-go/service/category"
	"PRACTICESTUFF/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
}