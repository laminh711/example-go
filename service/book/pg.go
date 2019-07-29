package book

import (
	"PRACTICESTUFF/example-go/domain"
	"context"

	"github.com/jinzhu/gorm"
)

// pgService implementer for BookService in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for BookService
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	return s.db.Create(p).Error
}

// Find implement Find for BookService
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p

	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement find all book for BookService
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	var res []domain.Book
	err := s.db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update implement update a book for BookService
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	err := s.db.Find(&old).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	old.Name = p.Name
	old.Author = p.Author
	old.Description = p.Description
	old.CategoryID = p.CategoryID

	return &old, s.db.Save(&old).Error
}

// Delete implement delete a book for BookService
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	err := s.db.Find(&old).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrRecordNotFound
		}
		return err
	}

	return s.db.Delete(old).Error
}

// IsCategoryExisted implement check category existence for BookService
func (s *pgService) IsCategoryExisted(ctx context.Context, cat *domain.Category) (bool, error) {
	category := domain.Category{Model: domain.Model{ID: cat.ID}}
	if err := s.db.Find(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
