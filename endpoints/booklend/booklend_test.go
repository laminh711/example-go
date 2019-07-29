package booklend

import (
	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
	booklendService "PRACTICESTUFF/example-go/service/booklend"
	"context"
	"testing"
	"time"
)

func TestMakeCreateEndpoint(t *testing.T) {
	mock := service.Service{
		BooklendService: &booklendService.ServiceMock{
			CreateFunc: func(_ context.Context, p *domain.Booklend) error {
				return nil
			},
		},
	}

	type args struct {
		req CreateRequest
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "update booklend endpoint parsed success",
			args: args{
				CreateRequest{
					CreateData{
						UserID: domain.MustGetUUIDFromString("cfa930f4-0f37-4d61-9314-5c2cb0993e44"),
						BookID: domain.MustGetUUIDFromString("a30cf9f4-f370-d641-9431-299cb03e5c44"),
						From:   time.Now().Local(),
						To:     time.Now().Local().Add(time.Hour * time.Duration(72)),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeCreateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			if err != nil {
				t.Fatalf("Create endpoint error %v", err)
			}
		})
	}
}
