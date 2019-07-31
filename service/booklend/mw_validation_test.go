// +build unit

package booklend

import (
	"PRACTICESTUFF/example-go/domain"
	"context"
	"net/http"
	"testing"
	"time"

	testutil "PRACTICESTUFF/example-go/config/database/pg/util"
)

func Test_validationMiddleware_Create(t *testing.T) {
	defaultCtx := context.Background()

	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	existedBook := domain.Book{}
	err = testDB.Create(&existedBook).Error
	if err != nil {
		t.Fatalf("Failed to create existedBook by error %v", err)
	}

	existedBorrowedBook := domain.Book{}
	err = testDB.Create(&existedBorrowedBook).Error
	if err != nil {
		t.Fatalf("Failed to create existedBorrowedBook by error %v", err)
	}

	existedUser := domain.User{}
	err = testDB.Create(&existedUser).Error
	if err != nil {
		t.Fatalf("Failed to create existedUser by error %v", err)
	}

	existedBooklend := domain.Booklend{
		BookID: existedBorrowedBook.ID,
		UserID: existedUser.ID,
		From:   time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
		To:     time.Date(2018, time.August, 10, 0, 0, 0, 0, time.UTC),
	}
	err = testDB.Create(&existedBooklend).Error
	if err != nil {
		t.Fatalf("Failed to create existedBooklend by error %v", err)
	}

	serviceMock := &ServiceMock{
		CreateFunc: func(ctx context.Context, p *domain.Booklend) error {
			return nil
		},
		IsBookExistedFunc: func(ctx context.Context, p *domain.Book) (bool, error) {
			return p.ID == existedBook.ID || p.ID == existedBorrowedBook.ID, nil
		},
		IsUserExistedFunc: func(ctx context.Context, p *domain.User) (bool, error) {
			return p.ID == existedUser.ID, nil
		},
		IsBooklendableFunc: func(ctx context.Context, p *domain.Booklend) (bool, error) {
			res := existedBooklend.BookID != p.BookID ||
				(existedBooklend.BookID == p.BookID && (existedBooklend.From.Sub(p.To) >= 0 || existedBooklend.To.Sub(p.From) <= 0))
			return res, nil
		},
	}

	fakeBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")
	fakeUserID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.Booklend
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "Pass validation cause of no booklend of this book existed",
			args: args{
				&domain.Booklend{
					BookID: existedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Pass validation cause of lending time not overlap",
			args: args{
				&domain.Booklend{
					BookID: existedBorrowedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Failed cause of invalid user id",
			args: args{
				&domain.Booklend{
					BookID: existedBook.ID,
					UserID: fakeUserID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrUserNotFound,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed cause of invalid book id",
			args: args{
				&domain.Booklend{
					BookID: fakeBookID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrBookNotFound,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed cause of invalid time span",
			args: args{
				&domain.Booklend{
					BookID: existedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 0, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrInvalidTimeSpan,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "book not lendable due to 'from' is overlapped with existed lending time",
			args: args{
				&domain.Booklend{
					BookID: existedBorrowedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 9, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 10, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrBookNotLendable,
			errorStatusCode: http.StatusConflict,
		},
		{
			name: "book not lendable due to 'to' is overlapped with existed lending time",
			args: args{
				&domain.Booklend{
					BookID: existedBorrowedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 5, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 6, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrBookNotLendable,
			errorStatusCode: http.StatusConflict,
		},
		{
			name: "book not lendable due to both 'from' and 'to' are overlapped with existed lending time",
			args: args{
				&domain.Booklend{
					BookID: existedBorrowedBook.ID,
					UserID: existedUser.ID,
					From:   time.Date(2018, time.August, 6, 0, 0, 0, 0, time.UTC),
					To:     time.Date(2018, time.August, 9, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr:         ErrBookNotLendable,
			errorStatusCode: http.StatusConflict,
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
