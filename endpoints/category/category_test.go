package category

import (
	"PRACTICESTUFF/example-go/domain"
	"PRACTICESTUFF/example-go/service"
	categoryService "PRACTICESTUFF/example-go/service/category"
	"context"
	"testing"
)

func TestMakeUpdateEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			UpdateFunc: func(_ context.Context, p *domain.Category) (*domain.Category, error) {
				return p, nil
			},
		},
	}

	type args struct {
		req UpdateRequest
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "update category endpoint parsed success",
			args: args{
				UpdateRequest{
					UpdateData{
						ID:   domain.MustGetUUIDFromString("415179ad-8067-4138-9b0d-41e0c68d4376"),
						Name: "Updated name",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeUpdateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			if err != nil {
				t.Fatalf("Update endpoint error %v", err)
			}
		})
	}
}

func TestMakeCreateEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			CreateFunc: func(_ context.Context, p *domain.Category) error {
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
			name: "update category endpoint parsed success",
			args: args{
				CreateRequest{
					CreateData{
						Name: "Created name",
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

func TestMakeFindCategoryEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			FindFunc: func(ctx context.Context, p *domain.Category) (*domain.Category, error) {
				return p, nil
			},
		},
	}

	type args struct {
		req FindRequest
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "find category endpoint parsed success",
			args: args{
				FindRequest{
					CategoryID: domain.MustGetUUIDFromString("cfa930f4-0f37-4d61-9314-5c2cb0993e44"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeFindEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			if err != nil {
				t.Fatalf("Find endpoint error %v", err)
			}
		})
	}
}

func TestMakeFindAllEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			FindAllFunc: func(_ context.Context) ([]domain.Category, error) {
				return []domain.Category{}, nil
			},
		},
	}
	type args struct {
		req FindAllRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "find all category endpoint parsed success",
			args: args{
				FindAllRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeFindAllEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeDeleteEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			DeleteFunc: func(_ context.Context, p *domain.Category) error {
				return nil
			},
		},
	}

	type args struct {
		req DeleteRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "delete category endpoint parsed success",
			args: args{
				DeleteRequest{
					CategoryID: domain.MustGetUUIDFromString("cfa930f4-0f37-4d61-9314-5c2cb0993e44"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeDeleteEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}