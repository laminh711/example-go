package category

import (
	"PRACTICESTUFF/example-go/domain"
	"context"
	"net/http"
	"reflect"
	"testing"

	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
)

func Test_validationMiddleware_Create(t *testing.T) {
	category := domain.Category{
		Name: "ExistedName",
	}
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.Category) error {
			return nil
		},
		IsNameDuplicateFunc: func(_ context.Context, nameToSearch string) (bool, error) {
			return category.Name == nameToSearch, nil
		},
	}

	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "Pass validation",
			args: args{
				&domain.Category{
					Name: "A category that longer than 5 characters",
				},
			},
		},
		{
			name: "Failed by empty name",
			args: args{
				&domain.Category{
					Name: "",
				},
			},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short name",
			args: args{
				&domain.Category{
					Name: "5char",
				},
			},
			wantErr:         ErrNameIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by duplicate name",
			args: args{
				&domain.Category{
					Name: "ExistedName",
				},
			},
			wantErr:         ErrNameIsDuplicated,
			errorStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			err := mw.Create(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}

				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}

			if tt.wantErr != nil {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	category := domain.Category{
		Name: "ExistedName",
	}
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Category) (*domain.Category, error) {
			return p, nil
		},
		IsNameDuplicateFunc: func(_ context.Context, nameToSearch string) (bool, error) {
			return category.Name == nameToSearch, nil
		},
	}

	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "Pass validation",
			args: args{
				&domain.Category{
					Name: "UpdateName",
				},
			},
		},
		{
			name: "Failed by empty name",
			args: args{
				&domain.Category{
					Name: "",
				},
			},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short name",
			args: args{
				&domain.Category{
					Name: "5char",
				},
			},
			wantErr:         ErrNameIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by duplicate name",
			args: args{
				&domain.Category{
					Name: "ExistedName",
				},
			},
			wantErr:         ErrNameIsDuplicated,
			errorStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			_, err := mw.Update(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}

				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}

			if tt.wantErr != nil {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_validationMiddleware_Find(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Category
		wantErr error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.Find(tt.args.ctx, tt.args.p)

			if err != nil {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Find() got %v, want %v", got, tt.want)
			}
			return
		})
	}
}

func Test_validationMiddleware_FindAll(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Category
		wantErr error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			got, err := mw.FindAll(tt.args.ctx)

			if err != nil {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Find() got %v, want %v", got, tt.want)
			}
			return
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.Delete(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
