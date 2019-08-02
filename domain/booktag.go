package domain

// BookTag book struct
type BookTag struct {
	Model
	BookID UUID `json:"book_id"`
	TagID  UUID `json:"tag_id"`
}

// Equal compare 2 BookTag
func (u BookTag) Equal(v BookTag) bool {
	return u.ID == v.ID && u.BookID == v.BookID && u.TagID == v.TagID
}
