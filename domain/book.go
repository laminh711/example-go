package domain

// Book book struct
type Book struct {
	Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CategoryID  UUID   `json:"category_id"`
}
