package booklend

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
)

// CreateData data for CreateBook
type CreateData struct {
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}

// CreateRequest request struct for CreateBook
type CreateRequest struct {
	Booklend []CreateData `json:"booklend"`
}

// CreateResponse response struct for CreateBook
type CreateResponse struct {
	Booklend []domain.Booklend `json:"booklend"`
}

// StatusCode customstatus code for success create Book
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Book
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(CreateRequest)
		// var (
		// 	booklend = &domain.Booklend{
		// 		UserID: req.Booklend.UserID,
		// 		BookID: req.Booklend.BookID,
		// 		From:   req.Booklend.From,
		// 		To:     req.Booklend.To,
		// 	}
		// )
		// err := s.BooklendService.Create(ctx, booklend)
		// if err != nil {
		// 	return nil, err
		// }

		res := CreateResponse{
			[]domain.Booklend{},
		}

		req := request.(CreateRequest)

		inpData := []domain.Booklend{}
		for _, cdata := range req.Booklend {
			booklend := &domain.Booklend{
				UserID: cdata.UserID,
				BookID: cdata.BookID,
				From:   cdata.From,
				To:     cdata.To,
			}
			inpData = append(inpData, *booklend)
		}

		err := s.BooklendService.CreateBatch(ctx, inpData)
		if err != nil {
			return nil, err
		}
		res.Booklend = inpData

		return res, nil

		// return CreateResponse{Booklend: *booklend}, nil
	}
}
