package endpoints

import (
	"PRACTICESTUFF/example-go/endpoints/book"
	"PRACTICESTUFF/example-go/service"

	"github.com/go-kit/kit/endpoint"

	"PRACTICESTUFF/example-go/endpoints/category"
	"PRACTICESTUFF/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser        endpoint.Endpoint
	FindAllUser     endpoint.Endpoint
	CreateUser      endpoint.Endpoint
	UpdateUser      endpoint.Endpoint
	DeleteUser      endpoint.Endpoint
	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	CreateBook  endpoint.Endpoint
	FindBook    endpoint.Endpoint
	FindAllBook endpoint.Endpoint
	UpdateBook  endpoint.Endpoint
	DeleteBook  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndpoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		CreateBook:  book.MakeCreateEndpoint(s),
		FindBook:    book.MakeFindEndpoint(s),
		FindAllBook: book.MakeFindAllEndpoint(s),
		UpdateBook:  book.MakeUpdateEndpoint(s),
		DeleteBook:  book.MakeDeleteEndpoint(s),
	}
}
