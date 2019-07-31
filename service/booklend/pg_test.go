// +build integration

package booklend

import (
	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
	"PRACTICESTUFF/example-go/domain"
	"context"
	"testing"
	"time"
)

func TestPGService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("failed to migrate tables by error %v", err)
	}

	type args struct {
		p *domain.Booklend
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "pgService create success",
			args: args{
				&domain.Booklend{
					BookID: domain.NewUUID(),
					UserID: domain.NewUUID(),
					From:   time.Now().Local(),
					To:     time.Now().Local().Add(time.Hour * time.Duration(8)),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pgService{
				db: testDB,
			}
			err := s.Create(context.Background(), tt.args.p)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestPGService_IsUserExisted(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	existedUser := domain.User{}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	type args struct {
		p *domain.User
	}

	fakeUserID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    bool
	}{
		{
			name: "Existed",
			args: args{
				p: &domain.User{
					Model: domain.Model{ID: existedUser.ID},
				},
			},
			want: true,
		},
		{
			name: "Not existed",
			args: args{
				p: &domain.User{
					Model: domain.Model{ID: fakeUserID},
				},
			},
			wantErr: true,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pgService{
				db: testDB,
			}

			got, err := s.IsUserExisted(context.Background(), tt.args.p)

			if err != nil {
				t.Errorf("pgService.IsUserExisted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("pgService.IsUserExisted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_IsBookExisted(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	existedBook := domain.Book{}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	type args struct {
		p *domain.Book
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    bool
	}{
		{
			name: "Existed",
			args: args{
				p: &domain.Book{
					Model: domain.Model{ID: existedBook.ID},
				},
			},
			want: true,
		},
		{
			name: "Not existed",
			args: args{
				p: &domain.Book{
					Model: domain.Model{ID: fakeBookID},
				},
			},
			wantErr: true,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pgService{
				db: testDB,
			}

			got, err := s.IsBookExisted(context.Background(), tt.args.p)

			if err != nil {
				t.Errorf("pgService.IsBookExisted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("pgService.IsBookExisted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGService_IsBooklendable(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	freeBook := domain.Book{}
	err = testDB.Create(&freeBook).Error
	if err != nil {
		t.Fatalf("Failed to create free book by error %v", err)
	}
	borrowedBook := domain.Book{}
	err = testDB.Create(&borrowedBook).Error
	if err != nil {
		t.Fatalf("Failed to create borrowed book by error %v", err)
	}
	borrowedBooklend := domain.Booklend{
		BookID: borrowedBook.ID,
		From:   time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
		To:     time.Date(2018, time.August, 10, 0, 0, 0, 0, time.UTC),
	}
	err = testDB.Create(&borrowedBooklend).Error
	if err != nil {
		t.Fatalf("Failed to create booklend by error %v", err)
	}

	type args struct {
		p *domain.Booklend
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    bool
	}{
		{
			name: "Lendable cause of no booklend of this book",
			args: args{
				p: &domain.Booklend{
					BookID: freeBook.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "Lendable cause of no time overlap",
			args: args{
				p: &domain.Booklend{
					BookID: borrowedBook.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "Not lendable because 'to' is in borrow time",
			args: args{
				p: &domain.Booklend{
					BookID: borrowedBook.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 6, 0, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
		{
			name: "Not lendable because 'from' is in borrow time",
			args: args{
				p: &domain.Booklend{
					BookID: borrowedBook.ID,
					From:   time.Date(2018, time.August, 9, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 11, 0, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
		{
			name: "Not lendable because 'from' and 'to' are in borrow time",
			args: args{
				p: &domain.Booklend{
					BookID: borrowedBook.ID,
					From:   time.Date(2018, time.August, 6, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 9, 0, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pgService{
				db: testDB,
			}

			got, err := s.IsBooklendable(context.Background(), tt.args.p)

			if err != nil {
				t.Errorf("pgService.IsBookExisted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("pgService.IsBookExisted() got = %v, want %v", got, tt.want)
			}
		})
	}
}
