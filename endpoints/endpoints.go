package endpoints

import (
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
	}
}