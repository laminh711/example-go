package domain

// Book book struct
type Book struct {
	Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CategoryID  UUID   `json:"category_id"`
}

// Equal check if another Book is equal to this (can't use DeepEqual due to weird Date Equal behavior)
func (u Book) Equal(v Book) bool {
	return u.ID == v.ID &&
		u.Name == v.Name &&
		u.Author == v.Author &&
		u.Description == v.Description &&
		u.CategoryID == v.CategoryID
}
