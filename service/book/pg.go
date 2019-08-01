package book

import (
	"PRACTICESTUFF/example-go/domain"
	"context"
	"fmt"
	"time"

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

// CreateBatch implement CreateBatch for BookService
func (s *pgService) CreateBatch(ctx context.Context, p []domain.Book) error {

	sqlValues := ""
	for i := 0; i < len(p); i++ {
		p[i].ID = domain.NewUUID()

		sqlValues += fmt.Sprintf("('%v', '%v', '%v', '%v', '%v')", p[i].ID, p[i].Name, p[i].Author, p[i].Description, p[i].CategoryID.String())
		if i == len(p)-1 {
			sqlValues += ";"
		} else {
			sqlValues += ",\n"
		}
	}

	t := time.Now().UTC()
	currentTimeValueOnPostgres := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%d+00:00",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000)
	returnValues := "SELECT * FROM books WHERE books.created_at <= " + "'" + currentTimeValueOnPostgres + "';"

	err := s.db.Raw("INSERT INTO books (id, name, author, description, category_id) VALUES " + sqlValues + returnValues).Scan(&p).Error

	return err
}

// Create implement Create for BookService
func (s *pgService) Create(ctx context.Context, p *domain.Book) error {
	return s.db.Create(p).Error
}

// Find implement Find for BookService
func (s *pgService) Find(ctx context.Context, p *domain.Book) (*domain.Book, error) {
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
func (s *pgService) FindAll(ctx context.Context, queries FindAllQueries) ([]domain.Book, error) {
	var res []domain.Book

	tmp := s.db.Where("books.name like ?", "%"+queries.Name+"%").Find(&res)
	err := tmp.Error
	if err != nil {
		return nil, err
	}

	if queries.Status != "" {
		unavailBookQuery := s.db.Table("booklends").
			Select("book_id").
			Where("booklends.from <= ? OR ? < booklends.from AND book_id IS NOT NULL", time.Now().Local(), time.Now().Local()).
			Joins("JOIN books ON books.id = booklends.book_id").
			QueryExpr()

		if queries.Status == "available" {
			tmp = tmp.Where("books.id NOT IN (?)", unavailBookQuery).Find(&res)
			err := tmp.Error
			if err != nil {
				return nil, err
			}
		}
		if queries.Status == "unavailable" {
			tmp = tmp.Where("books.id IN (?)", unavailBookQuery).Find(&res)
			err := tmp.Error
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update implement update a book for BookService
func (s *pgService) Update(ctx context.Context, p *domain.Book) (*domain.Book, error) {
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
func (s *pgService) Delete(ctx context.Context, p *domain.Book) error {
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
