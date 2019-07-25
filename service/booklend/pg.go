package booklend

import (
	"PRACTICESTUFF/example-go/domain"
	"context"

	"github.com/jinzhu/gorm"
)

// pgService implementer for Booklend service in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Booklend service
func (s *pgService) Create(_ context.Context, p *domain.Booklend) error {
	return s.db.Create(p).Error
}

// IsUserExisted implement IsUserExisted for Booklend service
func (s *pgService) IsUserExisted(_ context.Context, p *domain.User) (bool, error) {
	user := domain.User{Model: domain.Model{ID: p.ID}}

	if err := s.db.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsBookExisted implement IsBookExisted for Booklend service
func (s *pgService) IsBookExisted(_ context.Context, p *domain.Book) (bool, error) {
	book := domain.Book{Model: domain.Model{ID: p.ID}}

	if err := s.db.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsBookAvailable implement IsBookAvailable for Booklend service
func (s *pgService) IsBooklendable(ctx context.Context, p *domain.Booklend) (bool, error) {
	var res = domain.Booklend{}
	s.db.LogMode(true)
	temp := s.db.Where("(booklends.book_id = ?) AND ((booklends.from <= ? AND ? <= booklends.to) OR (booklends.from <= ? AND ? <= booklends.to))", p.BookID, p.From, p.From, p.To, p.To).Find(&res)
	err := temp.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
