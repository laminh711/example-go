// +build unit

package book

import (
	"PRACTICESTUFF/example-go/domain"
	"context"
	"net/http"
	"reflect"
	"testing"

	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
)

func Test_validationMiddleware_CreateBatch(t *testing.T) {
	existedCategory := domain.Category{
		Model: domain.Model{ID: domain.NewUUID()},
	}
	serviceMock := &ServiceMock{
		CreateBatchFunc: func(ctx context.Context, p []domain.Book) ([]domain.Book, error) {
			return []domain.Book{}, nil
		},
		IsCategoryExistedFunc: func(ctx context.Context, p *domain.Category) (bool, error) {
			return p.ID == existedCategory.ID, nil
		},
	}

	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	err = testDB.Create(&existedCategory).Error
	if err != nil {
		t.Fatalf("Failed to create existedCategory by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p []domain.Book
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
				[]domain.Book{
					domain.Book{
						Name:        "Call from angel",
						Description: "Detective-Romantic mix novel from Russo",
						CategoryID:  existedCategory.ID,
					},
				},
			},
		},
		{
			name: "Failed by empty name",
			args: args{
				[]domain.Book{
					domain.Book{
						Name:        "",
						Description: "A description for an empty book",
						CategoryID:  existedCategory.ID,
					},
				},
			},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short name",
			args: args{
				[]domain.Book{
					domain.Book{
						Name:        "5char",
						Description: "A description for a short book",
						CategoryID:  existedCategory.ID,
					},
				},
			},
			wantErr:         ErrNameIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by empty description",
			args: args{
				[]domain.Book{
					domain.Book{
						Name:        "You can be happy no matter what",
						Description: "",
						CategoryID:  existedCategory.ID,
					},
				},
			},
			wantErr:         ErrDescriptionIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short description",
			args: args{
				[]domain.Book{
					domain.Book{
						Name:        "Peter Pan",
						Description: "short",
						CategoryID:  existedCategory.ID,
					},
				},
			},
			wantErr:         ErrDescriptionIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by invalid category id",
			args: args{
				[]domain.Book{
					domain.Book{
						Name:        "book has fake category",
						Description: "A description for an book doesn't have valid categoryID",
						CategoryID:  fakeCategoryID,
					},
				},
			},
			wantErr:         ErrCategoryNotExisted,
			errorStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			_, err := mw.CreateBatch(defaultCtx, tt.args.p)
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

func Test_validationMiddleware_Create(t *testing.T) {
	existedCategory := domain.Category{
		Model: domain.Model{ID: domain.NewUUID()},
	}
	serviceMock := &ServiceMock{
		CreateFunc: func(ctx context.Context, p *domain.Book) error {
			return nil
		},
		IsCategoryExistedFunc: func(ctx context.Context, p *domain.Category) (bool, error) {
			return p.ID == existedCategory.ID, nil
		},
	}

	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	err = testDB.Create(&existedCategory).Error
	if err != nil {
		t.Fatalf("Failed to create existedCategory by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
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
				&domain.Book{
					Name:        "Call from angel",
					Description: "Detective-Romantic mix novel from Russo",
					CategoryID:  existedCategory.ID,
				},
			},
		},
		{
			name: "Failed by empty name",
			args: args{
				&domain.Book{
					Name:        "",
					Description: "A description for an empty book",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short name",
			args: args{
				&domain.Book{
					Name:        "5char",
					Description: "A description for a short book",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrNameIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by empty description",
			args: args{
				&domain.Book{
					Name:        "You can be happy no matter what",
					Description: "",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrDescriptionIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short description",
			args: args{
				&domain.Book{
					Name:        "Peter Pan",
					Description: "short",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrDescriptionIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by invalid category id",
			args: args{
				&domain.Book{
					Name:        "book has fake category",
					Description: "A description for an book doesn't have valid categoryID",
					CategoryID:  fakeCategoryID,
				},
			},
			wantErr:         ErrCategoryNotExisted,
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
	existedCategory := domain.Category{
		Model: domain.Model{ID: domain.NewUUID()},
	}
	serviceMock := &ServiceMock{
		UpdateFunc: func(ctx context.Context, p *domain.Book) (*domain.Book, error) {
			return p, nil
		},
		IsCategoryExistedFunc: func(ctx context.Context, p *domain.Category) (bool, error) {
			return p.ID == existedCategory.ID, nil
		},
	}

	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	err = testDB.Create(&existedCategory).Error
	if err != nil {
		t.Fatalf("Failed to create existedCategory by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
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
				&domain.Book{
					Name:        "a good name",
					Description: "a good description",
					CategoryID:  existedCategory.ID,
				},
			},
		},
		{
			name: "Failed by empty name",
			args: args{
				&domain.Book{
					Name:        "",
					Description: "A description for an empty name book",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short name",
			args: args{
				&domain.Book{
					Name:        "5char",
					Description: "A description for a short name book",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrNameIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by empty description",
			args: args{
				&domain.Book{
					Name:        "Book has no description",
					Description: "",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrDescriptionIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by short description",
			args: args{
				&domain.Book{
					Name:        "Book has short description",
					Description: "short",
					CategoryID:  existedCategory.ID,
				},
			},
			wantErr:         ErrDescriptionIsTooShort,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed by invalid category id",
			args: args{
				&domain.Book{
					Name:        "Book has invalid category id",
					Description: "A description for an book doesn't have valid categoryID",
					CategoryID:  fakeCategoryID,
				},
			},
			wantErr:         ErrCategoryNotExisted,
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
					t.Errorf("validationMiddleware.Update() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Update() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Update() error %v doesn't implement StatusCode()", err)
				}

				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Update() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}

			if tt.wantErr != nil {
				t.Errorf("validationMiddleware.Update() error = %v, wantErr = %v", err, tt.wantErr)
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
		p   *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Book
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

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	tag1 := domain.Tag{
		Name: "tag1",
	}
	err = testDB.Create(&tag1).Error
	if err != nil {
		t.Fatalf("Failed to create tag1 by error %v", err)
	}

	tag2 := domain.Tag{
		Name: "tag2",
	}
	err = testDB.Create(&tag2).Error
	if err != nil {
		t.Fatalf("Failed to create tag2 by error %v", err)
	}

	serviceMock := &ServiceMock{
		FindAllFunc: func(ctx context.Context, p FindAllQueries) ([]domain.Book, error) {
			return []domain.Book{}, nil
		},
		IsTagNameExistedFunc: func(ctx context.Context, t string) (bool, error) {
			return tag1.Name == t || tag2.Name == t, nil
		},
	}
	type args struct {
		q FindAllQueries
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "pass",
			args: args{
				FindAllQueries{
					Name:    "name is here",
					Status:  "available",
					TagName: "tag1",
				},
			},
		},
		{
			name: "failed by non existence tag name",
			args: args{
				FindAllQueries{
					TagName: "tag3",
				},
			},
			wantErr: ErrTagNameNotExisted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			_, err := mw.FindAll(context.Background(), tt.args.q)

			if err != nil {
				if err != tt.wantErr {
					t.Errorf("validationMiddleware.FindAll() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("validationMiddleware.FindAll() got %v, want %v", got, tt.want)
			// }
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
		p   *domain.Book
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
