// +build unit

package category

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
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				&domain.Category{
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

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Success update",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
					Name:  "Sports",
				},
			},
		},
		{
			name: "Failed update",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeCategoryID},
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

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}

	tests := []struct {
		name    string
		args    args
		want    *domain.Category
		wantErr error
	}{
		{
			name: "Success find correct user",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
				},
			},
			want: &domain.Category{
				Model: domain.Model{ID: category.ID},
			},
		},
		{
			name: "Failed find",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeCategoryID},
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

	category := domain.Category{}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	bookOfThisCat := domain.Book{
		CategoryID: category.ID,
	}
	err = testDB.Create(&bookOfThisCat).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}
	anotherBookOfThisCat := domain.Book{
		CategoryID: category.ID,
	}
	err = testDB.Create(&anotherBookOfThisCat).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	fakeCategoryID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Delete success",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: category.ID},
				},
			},
		},
		{
			name: "Fail delete due to wrong ID",
			args: args{
				&domain.Category{
					Model: domain.Model{ID: fakeCategoryID},
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

			// if delete success, need to check if all books that have this category also got deleted
			res := []domain.Book{}
			err = s.db.Find(&res).Error
			if err != nil {
				t.Fatalf("pgService.Delete() checking books after delete error %v", err)
			}

			if len(res) > 0 {
				t.Errorf("pgService.Delete() associated books not deleted after category deleted")
			}
		})
	}
}

func TestPGService_IsNameDuplicate(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	category := domain.Category{
		Name: "Duplicate",
	}
	err = testDB.Create(&category).Error
	if err != nil {
		t.Fatalf("Failed to create category by error %v", err)
	}

	type args struct {
		theName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "No duplicate",
			args: args{
				"NotDuplicate",
			},
			want: false,
		},
		{
			name: "Duplicate error",
			args: args{
				"Duplicate",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}

			got, err := s.IsNameDuplicate(context.Background(), tt.args.theName)

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
