// +build unit

package book

import (
	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
	"PRACTICESTUFF/example-go/domain"
	"context"
	"testing"
	"time"
)

func TestPGService_CreateBatch(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	type args struct {
		p []domain.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				[]domain.Book{
					domain.Book{
						Name: "Science-Fiction",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if _, err := s.CreateBatch(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

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

func EqualArrayResult(a []domain.Book, b []domain.Book) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !a[i].Equal(b[i]) {
			return false
		}
	}

	return true
}

func TestPGService_FindAll(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	book := domain.Book{
		Name: "book something",
	}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	borrowedBook := domain.Book{
		Name: "book borrowed",
	}
	err = testDB.Create(&borrowedBook).Error
	if err != nil {
		t.Fatalf("Failed to create borrowedBook by error %v", err)
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

	tag3 := domain.Tag{
		Name: "tag3",
	}
	err = testDB.Create(&tag3).Error
	if err != nil {
		t.Fatalf("Failed to create tag3 by error %v", err)
	}

	btag := domain.BookTag{
		BookID: book.ID,
		TagID:  tag1.ID,
	}
	err = testDB.Create(&btag).Error
	if err != nil {
		t.Fatalf("Failed to create btag by error %v", err)
	}

	btag2 := domain.BookTag{
		BookID: borrowedBook.ID,
		TagID:  tag2.ID,
	}
	err = testDB.Create(&btag2).Error
	if err != nil {
		t.Fatalf("Failed to create btag2 by error %v", err)
	}

	booklend := domain.Booklend{
		BookID: borrowedBook.ID,
		UserID: user.ID,
		From:   time.Now().Local().Add(time.Hour * time.Duration(-8)),
		To:     time.Now().Local().Add(time.Hour * time.Duration(8)),
	}
	err = testDB.Create(&booklend).Error
	if err != nil {
		t.Fatalf("Failed to create booklend by error %v", err)
	}

	type args struct {
		q FindAllQueries
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Book
		wantErr bool
	}{
		{
			name: "All, no queries",
			args: args{},
			want: []domain.Book{book, borrowedBook},
		},
		{
			name: "with name query",
			args: args{
				q: FindAllQueries{
					Name: "book",
				},
			},
			want: []domain.Book{book, borrowedBook},
		},
		{
			name: "with name query",
			args: args{
				q: FindAllQueries{
					Name: "book somethi",
				},
			},
			want: []domain.Book{book},
		},
		{
			name: "with name query",
			args: args{
				q: FindAllQueries{
					Name: "borrow",
				},
			},
			want: []domain.Book{borrowedBook},
		},
		{
			name: "with status query available",
			args: args{
				q: FindAllQueries{
					Status: "available",
				},
			},
			want: []domain.Book{book},
		},
		{
			name: "with status query unavailable",
			args: args{
				q: FindAllQueries{
					Status: "unavailable",
				},
			},
			want: []domain.Book{borrowedBook},
		},
		{
			name: "with tagname query tag1",
			args: args{
				q: FindAllQueries{
					TagName: "tag1",
				},
			},
			want: []domain.Book{book},
		},
		{
			name: "with tagname query tag2",
			args: args{
				q: FindAllQueries{
					TagName: "tag2",
				},
			},
			want: []domain.Book{borrowedBook},
		},
		{
			name: "with tagname query tag3",
			args: args{
				q: FindAllQueries{
					TagName: "tag3",
				},
			},
			want: []domain.Book{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.FindAll(context.Background(), tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !EqualArrayResult(got, tt.want) {
				t.Errorf("pgService.FindAll() \ngot %v, \nwant %v", got, tt.want)
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
		Name: "ExistedCategory",
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
