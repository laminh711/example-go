package domain

import "time"

// Booklend domain model for Booklend
type Booklend struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
