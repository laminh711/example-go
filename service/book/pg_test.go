package book

import (
	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
	"PRACTICESTUFF/example-go/domain"
	"context"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestPGService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Book{
					Name: "Science-Fiction",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_Update(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Success update",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: book.ID},
					Name:  "Sports",
				},
			},
		},
		{
			name: "Failed update",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: fakeBookID},
					Name:  "You shall not pass",
				},
			},
			wantErr: ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}

			_, err := s.Update(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Update() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_Find(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Book
		wantErr error
	}{
		{
			name: "Success find correct user",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: book.ID},
				},
			},
			want: &domain.Book{
				Model: domain.Model{ID: book.ID},
			},
		},
		{
			name: "Failed find",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: fakeBookID},
				},
			},
			wantErr: ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.Find(context.Background(), tt.args.p)

			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Find() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if got != nil && got.ID.String() != tt.want.ID.String() {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_Delete(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Delete success",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: book.ID},
				},
			},
		},
		{
			name: "Fail delete due to wrong ID",
			args: args{
				&domain.Book{
					Model: domain.Model{ID: fakeBookID},
				},
			},
			wantErr: ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			err := s.Delete(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Delete() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGService_IsCategoryExisted(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{
		Model: domain.Model{ID: domain.NewUUID()},
		Name:  "ExistedCategory",
	}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		category *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "Existed",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
				},
			},
			want: true,
		},
		{
			name: "Not existed",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeCategoryID},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}

			got, err := s.IsCategoryExisted(context.Background(), tt.args.category)

			if err != nil {
				if err != tt.wantErr {
					t.Errorf("pgService.IsNameDuplicate() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
			}

			if got != tt.want {
				t.Errorf("pgService.IsNameDuplicate() got %v, want %v", got, tt.want)
			}
		})
	}
}
